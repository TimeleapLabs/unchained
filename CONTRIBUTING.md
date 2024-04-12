# Contributing to Unchained

Thank you for your interest in contributing to Unchained! We value your
contributions and want to make sure the process is easy and beneficial for
everyone involved. Please take a moment to read these guidelines to ensure a
smooth contribution process.

## Contribution Guidelines

### 1. Contributor License Agreement (CLA)

Before your contributions can be accepted, you must sign a Contributor License
Agreement (CLA). This is a simple process that happens when you make your first
pull request (PR).

### 2. Forking and Branching Strategy

- **External Contributors**: Please fork the repository and submit your changes
  via a pull request from your fork.
- **Internal Contributors**: You should push your branches directly to the main
  project repository.

### 3. Issues and Pull Requests

- **Creating Issues**: Always create a GitHub issue before starting work on a
  pull request. Search for existing issues to avoid duplicates and check if
  someone else is already working on it.
- **Naming Branches**: Use semantic and meaningful branch names to make it clear
  what your contribution entails. Here are some good and bad examples:

  **Good Examples**

  - `add-{x}-to-{y}-{issue-number}`
  - `fix-{x}-in-{y}-{issue-number}`
  - `feature-{xyz}-{issue-number}`
  - `test-{xyz}-{issue-number}`

  **Bad Examples**

  - `bugfix`
  - `potatoes`
  - `mybranch`
  - `johns-first-pr`

### 4. Coding Standards

- **Documentation**: Your code must be well-documented and commented. Aim for
  clarity and simplicity; your code should be easily understandable without
  relying solely on comments.
- **Naming Conventions**: Use descriptive and appropriate names for variables,
  functions, and branches to clearly convey their purpose.
- **Pull Request Details**: Ensure your pull requests and issues clearly
  describe what you're trying to achieve. This helps reviewers understand your
  intentions and the impact of your contributions.

### 5. Pull Request Process

- Your PR should always target the `develop` branch.
- If the `develop` branch is updated before your PR is merged, you must rebase
  your commits to keep the history clean and avoid merge conflicts.

### 6. Commit Messages and Pull Requests

We use commitizen for commit and changelog management. You'll need the following
tools to get started:

- [Pre-commit](https://pre-commit.com/): Used for various pre-commit checks
  including commit message lint checks.
- [Commitizen](https://commitizen-tools.github.io): Used for creating commit
  messages and maintaining a logfile.
- [Golangci-lint](https://golangci-lint.run/): Used for Golang lint checks.

You'll need to run the following command to setup the pre-commit hooks:

```sh
pre-commit install --install-hooks --hook-type commit-msg --hook-type pre-push
```

## Conclusion

Following these guidelines helps us maintain a high standard of quality for our
project and makes the contribution process more efficient and effective. We look
forward to your contributions and are excited to see what we can achieve
together in Unchained!

Thank you for contributing!
