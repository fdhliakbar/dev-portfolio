@echo off
IF "%1"=="dev" (
    npm run dev
) ELSE IF "%1"=="start" (
    npm run start
) ELSE IF "%1"=="test" (
    npm run test
) ELSE (
    echo Unknown command: %1
    echo Available commands: dev, start, test
)
