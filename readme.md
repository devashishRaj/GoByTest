## NOTES on Test-driven development (TDD) with Go 
##### [Course link](https://github.com/quii/learn-go-with-tests)  

##### [Index](Index.md)

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Remember, keep trying to run the test and let the compiler guide you toward a solution.
- Write enough code to make the test pass
- Refactor

On the face of it this may seem tedious but sticking to the feedback loop is important.
Not only does it ensure that you have __relevant tests__, it helps ensure you __design good software__ by refactoring with the safety of tests.
__Seeing the test fail is an important check because it also lets you see what the error message looks like.__ As a developer it can be very hard to work with a codebase when failing tests do not give a clear idea as to what the problem is.
By ensuring your tests are fast and setting up your tools so that running tests is simple you can get in to a state of flow when writing your code.  

It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. __Every test has a cost.__  

__The definition of refactoring is that the code changes but the behaviour stays the same. If you have decided to do some refactoring in theory you should be able to make the commit without any test changes.__

### By not writing tests you are committing to manually checking your code by running your software which breaks your state of flow and you won't be saving yourself any time, especially in the long run.

```

Writing a test is just like writing a function, with a few rules
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- In order to use the *testing.T type, you need to import "testing"

```