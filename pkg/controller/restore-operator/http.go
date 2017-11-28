// Copyright 2017 The etcd-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	api "github.com/coreos/etcd-operator/pkg/apis/galera/v1alpha1"
	"github.com/coreos/etcd-operator/pkg/backup/backupapi"
	"github.com/coreos/etcd-operator/pkg/backup/reader"
	"github.com/coreos/etcd-operator/pkg/util/awsutil/s3factory"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	backupHTTPPath = backupapi.APIV1 + "/backup/"
	listenAddr     = "0.0.0.0:19999"
)

func (r *Restore) startHTTP() {
	http.HandleFunc(backupapi.APIV1+"/backup/", r.handleServeBackup)
	logrus.Infof("listening on %v", listenAddr)
	panic(http.ListenAndServe(listenAddr, nil))
}

func (r *Restore) handleServeBackup(w http.ResponseWriter, req *http.Request) {
	err := r.serveBackup(w, req)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// serveBackup parses incoming request url of the form /backup/<restore-name>
// get the etcd restore name.
// Then it returns the etcd cluster backup snapshot to the caller.
func (r *Restore) serveBackup(w http.ResponseWriter, req *http.Request) error {
	restoreName := string(req.URL.Path[len(backupHTTPPath):])
	if len(restoreName) == 0 {
		return errors.New("restore name is not specified")
	}

	obj := &api.EtcdRestore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      restoreName,
			Namespace: r.namespace,
		},
	}
	v, exists, err := r.indexer.Get(obj)
	if err != nil {
		return fmt.Errorf("failed to get restore CR for restore-name (%v): %v", restoreName, err)
	}
	if !exists {
		return fmt.Errorf("no restore CR found for restore-name (%v)", restoreName)
	}

	logrus.Infof("serving backup for restore CR %v", restoreName)
	cr := v.(*api.EtcdRestore)
	restoreSource := cr.Spec.RestoreSource
	var backupReader reader.Reader
	var path string

	switch {
	case restoreSource.S3 != nil:
		s3RestoreSource := restoreSource.S3
		if len(s3RestoreSource.AWSSecret) == 0 || len(s3RestoreSource.Path) == 0 {
			return errors.New("invalid s3 restore source field (spec.s3), must specify all required subfields")
		}

		s3Cli, err := s3factory.NewClientFromSecret(r.kubecli, r.namespace, s3RestoreSource.AWSSecret)
		if err != nil {
			return fmt.Errorf("failed to create S3 client: %v", err)
		}
		defer s3Cli.Close()

		backupReader = reader.NewS3Reader(s3Cli.S3)
		path = s3RestoreSource.Path
	default:
		return errors.New("restore CR must have a restore source specified")
	}

	rc, err := backupReader.Open(path)
	if err != nil {
		return fmt.Errorf("failed to read backup file(%v): %v", path, err)
	}
	defer rc.Close()

	_, err = io.Copy(w, rc)
	if err != nil {
		return fmt.Errorf("failed to write backup to %s: %v", req.RemoteAddr, err)
	}
	return nil
}
