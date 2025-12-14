// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

// AssertionTesterInterface defines an interface to be used for testing assertion methods.
type AssertionTesterInterface interface {
	TestMethod()
}

// AssertionTesterConformingObject is an object that conforms to the AssertionTesterInterface interface.
type AssertionTesterConformingObject struct{}

func (a *AssertionTesterConformingObject) TestMethod() {
}

// AssertionTesterNonConformingObject is an object that does not conform to the AssertionTesterInterface interface.
type AssertionTesterNonConformingObject struct{}
