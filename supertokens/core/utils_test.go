/*
 * Copyright (c) 2020, VRAI Labs and/or its affiliates. All rights reserved.
 *
 * This software is licensed under the Apache License, Version 2.0 (the
 * "License") as published by the Apache Software Foundation.
 *
 * You may not use this file except in compliance with the License. You may
 * obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package core

import (
	"testing"
)

func TestVersionCompare(t *testing.T) {
	if getLargestVersionFromIntersection([]string{"1.0"}, []string{"2.0"}) != nil {
		t.Error("should have failed")
	}

	if *getLargestVersionFromIntersection([]string{"1.0", "2.0"}, []string{"2.0"}) != "2.0" {
		t.Error("should have failed")
	}

	if *getLargestVersionFromIntersection([]string{"1.0.1", "1.1.19"}, []string{"2.0", "1.1.19", "2.1"}) != "1.1.19" {
		t.Error("should have failed")
	}

	if *getLargestVersionFromIntersection([]string{"1.0", "1.1.19", "1.1.2"}, []string{"1.0", "1.1.19", "1.1.2", "1.1.3"}) != "1.1.19" {
		t.Error("should have failed")
	}

	if *getLargestVersionFromIntersection([]string{"2.0", "1.1.19", "1.1.2"}, []string{"2.0", "1.1.19", "1.1.2", "1.1.3"}) != "2.0" {
		t.Error("should have failed")
	}
}
