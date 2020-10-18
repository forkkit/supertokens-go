# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.4.2] - 2020-09-19
### Fixed
- Fixed issue #13 - Do not clear cookies if they do not exist in the first place

## [1.4.1] - 2020-09-10
### Added
- Updates cookie setting to fix #4
- Adds `front-token` as per the latest FDI

## [1.4.0] - 2020-09-10
### Added
- Support for CDI 2.3 and FDI 1.2

## [1.3.0] - 2020-08-10
### Addition
- Adds `GetCORSAllowedHeaders` function to make CORS handling easier
- Compatibility with CDI 2.2

## [1.2.0] - 2020-07-02
### Addition
- Support for API key
- Compatibility with CDI 2.1

## [1.1.1] - 2020-06-28
### Bug fixes
- Sets cookie time based on seconds instead of milliseconds.

## [1.1.0] - 2020-06-16
### Changes
- Adds ability to set more config values like cookie domain, cookies paths etc.
- Adds ability to connect to https instance
- Changes how SuperTokens hosts are given

## [1.0.1] - 2020-06-14
### Fixed
- Converts frontendSDK info into a map[string]interface{} when returning it