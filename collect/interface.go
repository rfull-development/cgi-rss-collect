// Copyright (c) 2023 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package collect

// Raw is an interface for downloading raw data from a URL.
type Raw interface {
	Download() ([]byte, error)
}
