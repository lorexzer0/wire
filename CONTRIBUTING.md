# How to Contribute

We would love to accept your patches and contributions to this project. Here is
how you can help.

## Filing issues

Filing issues is an important way you can contribute to the Wire Project. We
want your feedback on things like bugs, desired API changes, or just anything
that isn't working for you.

### Bugs

If your issue is a bug, open one
[here](https://github.com/lorexzer0/wire/issues/new). The easiest way to file an
issue with all the right information is to run `go bug`. `go bug` will print out
a handy template of questions and system information that will help us get to
the root of the issue quicker.

### Changes

If you have a change you would like to see in Wire, please file an issue with
the necessary details.

## Contributing Code

We love accepting contributions! If your change is minor, please feel free
submit a [pull request](https://help.github.com/articles/about-pull-requests/).
If your change is larger, or adds a feature, please file an issue beforehand so
that we can discuss the change. You're welcome to file an implementation pull
request immediately as well, although we generally lean towards discussing the
change and then reviewing the implementation separately.

## Making a pull request

*   Follow the normal
    [pull request flow](https://help.github.com/articles/creating-a-pull-request/)
*   Test your changes using `go test ./...`. Please add tests that show the
    change does what it says it does, even if there wasn't a test in the first
    place.
*   Feel free to make as many commits as you want; we will squash them all into
    a single commit before merging your change.
*   Check the diffs, write a useful description (including something like
    `Fixes #123` if it's fixing a bug) and send the PR out.
*   Github will run tests against the PR. If a test fails, go back to the coding
    stage and try to fix the test and push the same branch again. You won't need
    to make a new pull request, the changes will be rolled directly into the PR
    you already opened.

## Code review

All submissions, including submissions by project members, require review. It is
almost never the case that a pull request is accepted without some changes
requested, so please do not be offended!

Once your PR is approved the reviewer will squash your commits into a single
commit, and then merge the commit onto the Wire main branch. Thank you!
