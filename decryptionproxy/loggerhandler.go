/**
 * Copyright 2020 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package decryptionproxy

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// LoggerHandler middleware adds logging before and after the main handler is invoked
func LoggerHandler(logger *logrus.Logger) Decorator {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				if logger.Level >= logrus.DebugLevel {
					fields := logrus.Fields{
						"method": r.Method,
						"rtt":    time.Since(start).String(),
						"path":   r.URL.RequestURI(),
					}
					entry := logger.WithFields(fields)
					entry.Debug("finished logging middleware")
				}
			}()
			handler.ServeHTTP(w, r)
		})
	}
}