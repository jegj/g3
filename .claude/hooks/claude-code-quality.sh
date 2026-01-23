#!/bin/bash

# Claude Code Hook: Code Quality Check
# Runs TypeScript check and lint fix after file modifications

echo "Running code quality checks..." >&2

# Run TypeScript check
echo "  Type checking..." >&2
npx tsc --noEmit 2>&1 >&2
if [ $? -ne 0 ]; then
  echo "TypeScript errors found" >&2
  echo '{"decision": "block", "reason": "TypeScript errors detected. Please fix type errors before proceeding."}'
  exit 0
fi

# Run lint fix
echo "  Linting and fixing..." >&2
npm run lint:fix 2>&1 >&2
if [ $? -ne 0 ]; then
  echo "Linting failed" >&2
  echo '{"decision": "block", "reason": "Linting failed. Please fix linting errors manually."}'
  exit 0
fi

echo "Code quality checks passed" >&2
echo '{"decision": "approve"}'
exit 0
