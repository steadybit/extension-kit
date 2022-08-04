// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

package util

func Ptr[T any](val T) *T {
	return &val
}
