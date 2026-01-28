---
title: "An approach to testing"
description: "How testify builds on top Go testing without replacing it"
weight: 20
---

{{% notice primary "TL;DR" "meteor" %}}
Software testing comes in two major styles: assertion-style (xUnit tradition: JUnit, pytest, NUnit, **testify**)
and BDD-style (RSpec, Cucumber, **Ginkgo**).

Both are valid approaches. However, we think that **testify's assertion-style naturally aligns with Go's core values**:
simplicity, explicitness, standard library first, minimal abstraction.

Testify brings powerful testing to Go developers who embrace these values:
zero-dependencies, reflection-based or generic assertions, and no framework (just works with `go test`).

**If you chose Go for its philosophy, assertion-style testing is the natural extension of those values to your test suite.**
{{% /notice %}}

---

## Make testing better. Keep it Go

**go-openapi/testify** follows a simple philosophy: **make Go testing better without reinventing it**.

Testify follows the assertion style: it is not a BDD framework. So you won't find chaining methods that produce
English-like sentences.

Unlike frameworks that introduce new paradigms and require specialized tooling,
testify builds directly on top of Go's great standard `testing` package.

It provides powerful assertions and utilities while preserving the familiar patterns that Go developers already know.
Testing patterns and constructs remain standard.

### Core Principles

**1. Zero Dependencies**

Testify has no external dependencies. Everything you need is self-contained, with internalized implementations of
required functionality. This means:

- No dependency conflicts in your project
- No supply chain security concerns
- No version compatibility issues
- Chrome is opt-in (all extra features that need additional dependencies are opt-in)

**2. Standard Go Compatibility**

Works seamlessly with `go test` and the standard library:

- No special CLI tools required
- No framework-specific test runners
- Standard Go subtests with `t.Run()`
- Native IDE support out of the box
- Works with any Go test runner

**3. Type Safety with Generics**

Testify embraces Go's type system:

- Most assertions come with a **generic variant** for compile-time type safety
- Catch type mismatches before tests even run
- On average **10x faster** than reflection-based assertions
- Full type inference: no manual type parameters needed
- Complex cases that require dynamic typing use go reflection

**4. Simplicity and Clarity**

Keep testing straightforward:

- Function-based assertions with clear semantics
- No new DSL to learn
- Minimal cognitive overhead
- Immediate productivity for any Go developer

---

## Testing Styles: Assertion vs. BDD

Software testing has evolved into two primary styles, each with passionate advocates across programming communities.

### Assertion-Style Testing (xUnit tradition)

**Core idea**: Write tests as regular code with explicit assertions.

Originating with Kent Beck's SUnit (Smalltalk) and popularized by JUnit (Java), this style emphasizes:

- Tests are functions/methods in the language
- Direct assertion calls verify behavior
- Standard language constructs for organization
- Minimal framework abstraction

**Examples across languages:**

{{< tabs >}}
{{% tab title="JUnit (Java)" %}}
```java
// JUnit (Java)
@Test
public void testUserCreation() {
    User user = createUser("alice@example.com");
    assertNotNull(user);
    assertEquals("alice@example.com", user.getEmail());
}
```
{{% /tab %}}

{{% tab title="pytest (python)" %}}
```python
# pytest (Python)
def test_user_creation():
    user = create_user("alice@example.com")
    assert user is not None
    assert user.email == "alice@example.com"
```
{{% /tab %}}

{{% tab title="NUnit (C#)" %}}
```csharp
// NUnit (C#)
[Test]
public void TestUserCreation()
{
    var user = CreateUser("alice@example.com");
    Assert.IsNotNull(user);
    Assert.AreEqual("alice@example.com", user.Email);
}
```
{{% /tab %}}

{{% tab title="Testify (go)" %}}
```go
// Testify (go)
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestUserCreation(t *testing.T) {
    user := CreateUser("alice@example.com")
    require.NotNil(t, user)
    assert.Equal(t, "alice@example.com", user.Email)
}
```
{{% /tab %}}
{{% /tabs %}}

**Frameworks**: JUnit, NUnit, xUnit.net, pytest, PHPUnit, Go's `testing` package... and `testify`.

### BDD-Style Testing (Behavior-Driven Development)

**Core idea**: Write tests as executable specifications in narrative form.

Originating with RSpec (Ruby) and influenced by Dan North's BDD methodology, this style emphasizes:

- Tests describe behavior in natural language structure
- Hierarchical organization (describe/context/it)
- Focus on readability and documentation value
- Framework-specific DSL

**Examples across languages:**

{{< tabs >}}
{{% tab title="RSpec (Ruby)" %}}
```ruby
# RSpec (Ruby)
describe "User creation" do
  it "creates a valid user" do
    user = create_user("alice@example.com")
    expect(user).not_to be_nil
    expect(user.email).to eq("alice@example.com")
  end
end
```
{{% /tab %}}

{{% tab title="Jasmine (JS)" %}}
```javascript
// Jasmine/Mocha (JavaScript)
describe("User creation", function() {
  it("creates a valid user", function() {
    const user = createUser("alice@example.com");
    expect(user).not.toBe(null);
    expect(user.email).toEqual("alice@example.com");
  });
});
```
{{% /tab %}}

{{% tab title="behave (python)" %}}
```python
# behave (Python)
Scenario: User creation
  Given a valid email address
  When I create a user with "alice@example.com"
  Then the user should exist
  And the email should be "alice@example.com"
```
{{% /tab %}}

{{% tab title="Ginkgo (go)" %}}
```go
// Ginkgo
import (
    "github.com/onsi/ginkgo/v2"
    "github.com/onsi/gomega"
)

var _ = ginkgo.Describe("User creation", func() {
            ginkgo.It("creates a valid user", func() {
                user := CreateUser("alice@example.com")
                gomega.Expect(user).NotTo(gomega.BeNil())
                gomega.Expect(user.Email).To(gomega.Equal("alice@example.com"))
            })
        })
```
{{% /tab %}}
{{% /tabs %}}

**Frameworks**: RSpec, Jasmine, Mocha, Cucumber, behave, **Ginkgo/Gomega**

### Both Are Valid

**Assertion-style strengths:**

- Low cognitive overhead (just code)
- Minimal framework abstraction
- IDE tooling works naturally
- Easy to learn and adopt

**BDD-style strengths:**

- Readable test specifications
- Natural hierarchical organization
- Self-documenting intent
- Stakeholder-friendly output

**The debate continues** across all programming communities. Neither style is objectively superior; they optimize for
different values and team preferences.

---

## Assertion-Style and Go Values

While both styles have merit in general, **assertion-style testing aligns naturally with Go's core philosophy**.

### Go's Design Values

Go emphasizes:

- **Simplicity**: maximize clarity
- **Explicitness**: No magic, no hidden behavior
- **Standard library first**: Build on solid foundations
- **Readability**: Code is read more than written
- **Minimal abstraction**: Minimize concepts

### How Assertion-Style Matches Go

**1. Simplicity**

Assertion-style keeps tests simple: they're just Go functions.

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result) // Straightforward, even though we haven't written "When().Two().Plus().Three().IsNot(5).Fail()"...
}
```

No new mental model. No framework semantics to learn. If you know Go, you know how to test with our lib.

**2. Explicitness**

Every assertion is an explicit function call with clear semantics:

```go
assert.NotNil(t, user)                     // Explicit: check for nil
assert.ErrorIs(t, err, ErrNotFound)        // Explicit: check error identity
assert.ElementsMatch(t, expected, actual)  // Explicit: check collection equality
```

Compare to matcher-based approaches where behavior is composed through framework objects.
Assertion-style makes test intent immediately clear _to a programmer_.

**3. Standard Library First**

Testify builds on `testing.T`: no replacement, just enhancement.

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestUserAPI(t *testing.T) {
	t.Run("creation", func(t *testing.T) { // Standard Go subtest
		t.Parallel() // Standard Go test parallelism

		user := CreateUser("alice@example.com")
		assert.NotNil(t, user) // Enhanced with assertions
	})
}
```

Works with `go test`. Works with standard tooling. Works with the Go ecosystem.

**4. Readability Through Directness**

Go prioritizes code that's easy to read and understand:

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

// Clear, direct, readable
func TestEmailValidation(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"alice@example.com", true},
		{"invalid-email", false},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if tt.valid {
				assert.NoError(t, err)

                return
			}

			assert.Error(t, err)
		})
	}
}
```

Standard control flow. Standard Go idioms. No DSL to decode.
Better for most developers, perhaps less so for stakeholders not familiar with Go.

**5. Minimal Abstraction**

Go avoids abstraction for abstraction's sake. Testify provides assertions: nothing more.

- No test lifecycle framework
- No dependency injection system
- No specialized runners
- No mandatory patterns

Just functions that verify behavior and produce clear errors.
Solve the testing problem, don't create a testing ecosystem.

### The Natural Fit

If you appreciate Go's philosophy and if you choose Go because you value simplicity, explicitness, and building on
standards, then **assertion-style testing is the natural extension of those values to your test suite**.

BDD frameworks serve teams with different priorities (narrative specifications, framework-managed workflows, stakeholder
communication). Those are valid priorities. But they optimize for values orthogonal to Go's design philosophy.
For such teams, **Ginkgo/Gomega** provides a great BDD testing framework.

**For Go developers who embrace Go values, assertion-style testing is the idiomatic approach.**. And **testify** is the tool.

---

## Assertion-Style in Go: Testify vs. BDD Frameworks

Go's testing ecosystem reflects the broader assertion-vs-BDD divide:

### Different Philosophies

| Aspect | Testify (Assertion-Style) | Ginkgo/Gomega (BDD-Style) |
|--------|---------------------------|---------------------------|
| **Testing style** | xUnit tradition | BDD tradition |
| **Approach** | Enhance standard testing | Replace with BDD framework |
| **Integration** | Works with `go test` directly | Requires `ginkgo` CLI tool |
| **Learning curve** | Immediate (standard Go) | Moderate (new DSL) |
| **Dependencies** | Zero external packages | Multiple framework packages |
| **Type safety** | 40 generic assertions | Reflection-based matchers |
| **Organization** | Standard Go subtests | Narrative hierarchy (Describe/Context/It) |
| **Go philosophy** | Aligns with Go values | Different priorities |

### Example Comparison

**Assertion-style (Testify):**

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestUserCreation(t *testing.T) {
	t.Parallel()

	user := CreateUser("alice@example.com")

	// Clear, type-safe assertions
	require.NotNil(t, user)
	assert.EqualT(t, "alice@example.com", user.Email) // Compile-time type check
	assert.True(t, user.Active)
}
```

**BDD-style (Ginkgo/Gomega):**

```go
var _ = Describe("User Creation", func() {
	var user *User

	BeforeEach(func() {
		user = CreateUser("alice@example.com")
	})

	It("creates a valid user", func() {
		Expect(user).ToNot(BeNil())
		Expect(user.Email).To(Equal("alice@example.com"))
		Expect(user.Active).To(BeTrue())
	})
})
```

Both approaches are valid. They reflect different testing philosophies that span the entire software industry. The
question for Go developers is: **which style aligns with the values that drew you to Go in the first place?**

---

## See also

- [API Reference](../api/) - Browse all 120 assertions by domain
- [Generics Guide](../usage/GENERICS.md) - Leverage type-safe assertions
- [Migration Guide](../usage/MIGRATION.md) - Switch from stretchr/testify
- [Examples](../usage/EXAMPLES.md) - See testify in action

---
