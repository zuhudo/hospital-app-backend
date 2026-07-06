# Contributing to Hospital App

Thank you for your interest in contributing! 🎉

## Getting Started

1. **Fork** the repository
2. **Clone** your fork: `git clone https://github.com/YOUR_USERNAME/REPO_NAME.git`
3. **Create** a branch: `git checkout -b feature/amazing-feature`
4. **Make** your changes
5. **Test** your changes
6. **Commit**: `git commit -m "feat: add amazing feature"`
7. **Push**: `git push origin feature/amazing-feature`
8. **Open** a Pull Request

## Development Setup

### Prerequisites

- **Flutter SDK** 3.2+ (for mobile)
- **Go** 1.21+ (for backend)
- **Node.js** 18+ (for web/dashboard)
- **pnpm** (for web/dashboard)

### Running Locally

```bash
# Backend
cd backend
cp .env.example .env
go run main.go

# Flutter App
cd app
flutter pub get
flutter run

# Public Website
cd web/public
pnpm install
pnpm dev

# Admin Dashboard
cd web/dashboard
pnpm install
pnpm dev
```

## Commit Convention

We use [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` — New feature
- `fix:` — Bug fix
- `docs:` — Documentation
- `style:` — Formatting, missing semicolons, etc.
- `refactor:` — Code restructuring without behavior change
- `test:` — Adding tests
- `chore:` — Maintenance tasks

## Pull Request Guidelines

- Fill in the PR template completely
- Link related issues
- Include screenshots for UI changes
- Ensure all CI checks pass
- Request review from maintainers

## Code Style

- **Dart**: Follow [Effective Dart](https://dart.dev/guides/language/effective-dart)
- **Go**: Follow [Effective Go](https://go.dev/doc/effective_go) and run `gofmt`
- **TypeScript**: Follow the project's ESLint/Prettier config

## Reporting Issues

Use the issue templates provided. Include:
- Clear description
- Steps to reproduce
- Expected vs actual behavior
- Environment details

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
