# Imversed Documentation

All commands described here should be executed inside `docs` directory.

## Prerequisites

Install required Node.js modules:

```bash
npm i
```

Copy documentation files located outside `docs` directory:

```bash
npm run prepare
```

## Develop

Start development server, perform any changes and navigate to http://localhost:8080 to see the changes:

```bash
npm run dev
```

## Publish

Make sure that `md`-files located outside `docs` directory are up to date:

```bash
npm run prepare
```

Build static files (will be placed in `docs/src/.vuepress/dist`):

```bash
npm run build
```

Login to Firebase if not yet and deploy the statics:

```bash
firebase login
firebase deploy
```
