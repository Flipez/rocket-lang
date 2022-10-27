# Changelog

## [Unreleased](https://github.com/flipez/rocket-lang/tree/HEAD)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.18.0...HEAD)

**Implemented enhancements:**

- Add missing math functions [\#127](https://github.com/Flipez/rocket-lang/issues/127)
- Document `else if`  [\#121](https://github.com/Flipez/rocket-lang/issues/121)
- Triginometric functions and pi [\#104](https://github.com/Flipez/rocket-lang/issues/104)
- \[foreach/hash\] Swap key/value order [\#132](https://github.com/Flipez/rocket-lang/pull/132) ([Flipez](https://github.com/Flipez))
- allow module name to be defined optionally [\#123](https://github.com/Flipez/rocket-lang/pull/123) ([Flipez](https://github.com/Flipez))
- introduce an object iterator [\#122](https://github.com/Flipez/rocket-lang/pull/122) ([MarkusFreitag](https://github.com/MarkusFreitag))
- add Time builtin [\#113](https://github.com/Flipez/rocket-lang/pull/113) ([Flipez](https://github.com/Flipez))
- move exit and raise to os module [\#111](https://github.com/Flipez/rocket-lang/pull/111) ([Flipez](https://github.com/Flipez))
- move open into IO module [\#109](https://github.com/Flipez/rocket-lang/pull/109) ([Flipez](https://github.com/Flipez))
- Implement `and`, `or` , `&&` and `||` [\#102](https://github.com/Flipez/rocket-lang/pull/102) ([Flipez](https://github.com/Flipez))

**Fixed bugs:**

- Nested loop seems to be broken [\#120](https://github.com/Flipez/rocket-lang/issues/120)
- Line number for error is off [\#114](https://github.com/Flipez/rocket-lang/issues/114)
- \[stdlib/math\] Seed rand on each call [\#131](https://github.com/Flipez/rocket-lang/pull/131) ([Flipez](https://github.com/Flipez))
- \[stdlib/math\] fix typo in docs [\#130](https://github.com/Flipez/rocket-lang/pull/130) ([RaphaelPour](https://github.com/RaphaelPour))
- document else if [\#124](https://github.com/Flipez/rocket-lang/pull/124) ([Flipez](https://github.com/Flipez))
- Fix line calculation in error messages [\#117](https://github.com/Flipez/rocket-lang/pull/117) ([Flipez](https://github.com/Flipez))
- fix repl cmd handling [\#107](https://github.com/Flipez/rocket-lang/pull/107) ([Flipez](https://github.com/Flipez))

**Merged pull requests:**

- Hash convenience set [\#110](https://github.com/Flipez/rocket-lang/pull/110) ([MarkusFreitag](https://github.com/MarkusFreitag))
- change builtin implementation [\#108](https://github.com/Flipez/rocket-lang/pull/108) ([Flipez](https://github.com/Flipez))
- use inter font [\#103](https://github.com/Flipez/rocket-lang/pull/103) ([Flipez](https://github.com/Flipez))
- link playground [\#101](https://github.com/Flipez/rocket-lang/pull/101) ([Flipez](https://github.com/Flipez))
- Add Playground [\#100](https://github.com/Flipez/rocket-lang/pull/100) ([Flipez](https://github.com/Flipez))
- add release posts; fix favicon [\#99](https://github.com/Flipez/rocket-lang/pull/99) ([Flipez](https://github.com/Flipez))
- add first release post [\#98](https://github.com/Flipez/rocket-lang/pull/98) ([Flipez](https://github.com/Flipez))
- Update docs page [\#97](https://github.com/Flipez/rocket-lang/pull/97) ([Flipez](https://github.com/Flipez))

## [v0.18.0](https://github.com/flipez/rocket-lang/tree/v0.18.0) (2022-07-25)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.17.1...v0.18.0)

**Implemented enhancements:**

- Parser: add more errors for bad syntax [\#68](https://github.com/Flipez/rocket-lang/issues/68)

**Merged pull requests:**

- Add support for single quotes and escape double quotes [\#96](https://github.com/Flipez/rocket-lang/pull/96) ([Flipez](https://github.com/Flipez))
- add json object [\#95](https://github.com/Flipez/rocket-lang/pull/95) ([Flipez](https://github.com/Flipez))
- remove args from next and break [\#94](https://github.com/Flipez/rocket-lang/pull/94) ([Flipez](https://github.com/Flipez))

## [v0.17.1](https://github.com/flipez/rocket-lang/tree/v0.17.1) (2022-07-03)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.17.0...v0.17.1)

**Closed issues:**

- Add support for `mosel`-loop [\#85](https://github.com/Flipez/rocket-lang/issues/85)
- Collect Emoji Syntax Ideas [\#9](https://github.com/Flipez/rocket-lang/issues/9)

**Merged pull requests:**

- Foreach improvements [\#92](https://github.com/Flipez/rocket-lang/pull/92) ([Flipez](https://github.com/Flipez))

## [v0.17.0](https://github.com/flipez/rocket-lang/tree/v0.17.0) (2022-07-03)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.16.0...v0.17.0)

**Merged pull requests:**

- Add nil and replace null with it [\#91](https://github.com/Flipez/rocket-lang/pull/91) ([Flipez](https://github.com/Flipez))

## [v0.16.0](https://github.com/flipez/rocket-lang/tree/v0.16.0) (2022-07-02)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.15.1...v0.16.0)

**Implemented enhancements:**

- Fix indexing [\#79](https://github.com/Flipez/rocket-lang/issues/79)
- Improve file.write\(\) [\#37](https://github.com/Flipez/rocket-lang/issues/37)
- add basic http response [\#84](https://github.com/Flipez/rocket-lang/pull/84) ([Flipez](https://github.com/Flipez))
- add .to\_json\(\) [\#82](https://github.com/Flipez/rocket-lang/pull/82) ([Flipez](https://github.com/Flipez))
- add networking [\#81](https://github.com/Flipez/rocket-lang/pull/81) ([Flipez](https://github.com/Flipez))
- add ast tests [\#77](https://github.com/Flipez/rocket-lang/pull/77) ([Flipez](https://github.com/Flipez))
- Implement while loop [\#75](https://github.com/Flipez/rocket-lang/pull/75) ([MarkusFreitag](https://github.com/MarkusFreitag))
- add ternary [\#73](https://github.com/Flipez/rocket-lang/pull/73) ([Flipez](https://github.com/Flipez))
- object/integer: satisfy iterable interface [\#66](https://github.com/Flipez/rocket-lang/pull/66) ([RaphaelPour](https://github.com/RaphaelPour))
- Improve Infix operator [\#65](https://github.com/Flipez/rocket-lang/pull/65) ([Flipez](https://github.com/Flipez))

**Fixed bugs:**

- Fix indexable object types [\#80](https://github.com/Flipez/rocket-lang/pull/80) ([Flipez](https://github.com/Flipez))

**Merged pull requests:**

- Add support for `next()` and `break()` [\#90](https://github.com/Flipez/rocket-lang/pull/90) ([Flipez](https://github.com/Flipez))
- Remove support for curly braces in foreach, while, if and function [\#89](https://github.com/Flipez/rocket-lang/pull/89) ([Flipez](https://github.com/Flipez))
- Return bytesWritten instead of True on successfull file.write\(\) [\#88](https://github.com/Flipez/rocket-lang/pull/88) ([Flipez](https://github.com/Flipez))
- Apt repo [\#87](https://github.com/Flipez/rocket-lang/pull/87) ([Flipez](https://github.com/Flipez))
- add tests [\#78](https://github.com/Flipez/rocket-lang/pull/78) ([Flipez](https://github.com/Flipez))
- add missing quotation mark in string documentation [\#76](https://github.com/Flipez/rocket-lang/pull/76) ([Tch1b0](https://github.com/Tch1b0))
- float: adds plz\_i object function [\#72](https://github.com/Flipez/rocket-lang/pull/72) ([RaphaelPour](https://github.com/RaphaelPour))
- add modulo [\#70](https://github.com/Flipez/rocket-lang/pull/70) ([Flipez](https://github.com/Flipez))

## [v0.15.1](https://github.com/flipez/rocket-lang/tree/v0.15.1) (2022-01-22)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.15.0...v0.15.1)

**Merged pull requests:**

- build more packages [\#63](https://github.com/Flipez/rocket-lang/pull/63) ([Flipez](https://github.com/Flipez))

## [v0.15.0](https://github.com/flipez/rocket-lang/tree/v0.15.0) (2022-01-21)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.14.2...v0.15.0)

**Implemented enhancements:**

- Feature Request: Add CLI args / methods [\#52](https://github.com/Flipez/rocket-lang/issues/52)
- Foreach error abort [\#62](https://github.com/Flipez/rocket-lang/pull/62) ([Flipez](https://github.com/Flipez))
- add cli flags [\#57](https://github.com/Flipez/rocket-lang/pull/57) ([Flipez](https://github.com/Flipez))
- Improve Error Messages [\#56](https://github.com/Flipez/rocket-lang/pull/56) ([Flipez](https://github.com/Flipez))

**Fixed bugs:**

- Foreach should abort on error and return [\#50](https://github.com/Flipez/rocket-lang/issues/50)

**Merged pull requests:**

- examples: Update aoc/2015/day1 to 0.14.1 [\#55](https://github.com/Flipez/rocket-lang/pull/55) ([Kjarrigan](https://github.com/Kjarrigan))

## [v0.14.2](https://github.com/flipez/rocket-lang/tree/v0.14.2) (2022-01-20)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.14.1...v0.14.2)

**Fixed bugs:**

- Bug: open\(\) with two arguments does not work [\#54](https://github.com/Flipez/rocket-lang/issues/54)
- Bug: Comments are broken [\#53](https://github.com/Flipez/rocket-lang/issues/53)
- fix comments [\#61](https://github.com/Flipez/rocket-lang/pull/61) ([Flipez](https://github.com/Flipez))

**Merged pull requests:**

- Improve docs [\#60](https://github.com/Flipez/rocket-lang/pull/60) ([Flipez](https://github.com/Flipez))
- Improve docs [\#59](https://github.com/Flipez/rocket-lang/pull/59) ([Flipez](https://github.com/Flipez))
- Improve docs [\#58](https://github.com/Flipez/rocket-lang/pull/58) ([Flipez](https://github.com/Flipez))

## [v0.14.1](https://github.com/flipez/rocket-lang/tree/v0.14.1) (2022-01-18)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.14.0...v0.14.1)

## [v0.14.0](https://github.com/flipez/rocket-lang/tree/v0.14.0) (2022-01-18)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.13.0...v0.14.0)

**Implemented enhancements:**

- unify object creation with New methods [\#51](https://github.com/Flipez/rocket-lang/pull/51) ([MarkusFreitag](https://github.com/MarkusFreitag))
- improve tests [\#49](https://github.com/Flipez/rocket-lang/pull/49) ([Flipez](https://github.com/Flipez))
- implement LT\_EQ and GT\_EQ [\#48](https://github.com/Flipez/rocket-lang/pull/48) ([MarkusFreitag](https://github.com/MarkusFreitag))
- clean up lexer [\#47](https://github.com/Flipez/rocket-lang/pull/47) ([Flipez](https://github.com/Flipez))
- Restructure Code [\#45](https://github.com/Flipez/rocket-lang/pull/45) ([Flipez](https://github.com/Flipez))

## [v0.13.0](https://github.com/flipez/rocket-lang/tree/v0.13.0) (2022-01-16)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.12.0...v0.13.0)

**Implemented enhancements:**

- Improve `if` Statement [\#41](https://github.com/Flipez/rocket-lang/issues/41)
- Remove `let` as it is deprecated [\#38](https://github.com/Flipez/rocket-lang/issues/38)
- implement type float [\#43](https://github.com/Flipez/rocket-lang/pull/43) ([Flipez](https://github.com/Flipez))
- make curly braces optional [\#42](https://github.com/Flipez/rocket-lang/pull/42) ([Flipez](https://github.com/Flipez))
- replace fn identifier with def [\#40](https://github.com/Flipez/rocket-lang/pull/40) ([Flipez](https://github.com/Flipez))
- remove let [\#39](https://github.com/Flipez/rocket-lang/pull/39) ([Flipez](https://github.com/Flipez))

## [v0.12.0](https://github.com/flipez/rocket-lang/tree/v0.12.0) (2022-01-15)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.11.1...v0.12.0)

**Implemented enhancements:**

- Improve file handling [\#35](https://github.com/Flipez/rocket-lang/issues/35)
- rewrite file handle, adjust file.seek\(\) [\#36](https://github.com/Flipez/rocket-lang/pull/36) ([Flipez](https://github.com/Flipez))

**Fixed bugs:**

- file.read\(\) and file.line\(\) return unexpected results [\#33](https://github.com/Flipez/rocket-lang/issues/33)

## [v0.11.1](https://github.com/flipez/rocket-lang/tree/v0.11.1) (2022-01-14)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.11.0...v0.11.1)

**Fixed bugs:**

- rewrite file.read\(\) and file.lines\(\) [\#34](https://github.com/Flipez/rocket-lang/pull/34) ([Flipez](https://github.com/Flipez))

## [v0.11.0](https://github.com/flipez/rocket-lang/tree/v0.11.0) (2022-01-13)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.10.0...v0.11.0)

**Implemented enhancements:**

- Document builtins [\#26](https://github.com/Flipez/rocket-lang/issues/26)
- Allow to load code from other scripts [\#10](https://github.com/Flipez/rocket-lang/issues/10)
- move builtins to stdlib layout [\#28](https://github.com/Flipez/rocket-lang/pull/28) ([Flipez](https://github.com/Flipez))

**Merged pull requests:**

- Release 0.11 [\#30](https://github.com/Flipez/rocket-lang/pull/30) ([Flipez](https://github.com/Flipez))

## [v0.10.0](https://github.com/flipez/rocket-lang/tree/v0.10.0) (2021-12-27)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.9.7...v0.10.0)

**Implemented enhancements:**

- Avoid using append when creating arrays [\#1](https://github.com/Flipez/rocket-lang/issues/1)
- Add support for loops [\#4](https://github.com/Flipez/rocket-lang/issues/4)

**Fixed bugs:**

- Omit the `null` on execution ending [\#6](https://github.com/Flipez/rocket-lang/issues/6)

**Closed issues:**

- Weird variable scope [\#13](https://github.com/Flipez/rocket-lang/issues/13)
- Add version-string to release-filenames [\#12](https://github.com/Flipez/rocket-lang/issues/12)

**Merged pull requests:**

- sourround strings in inspect with quotes [\#25](https://github.com/Flipez/rocket-lang/pull/25) ([Flipez](https://github.com/Flipez))
- Docs [\#24](https://github.com/Flipez/rocket-lang/pull/24) ([Flipez](https://github.com/Flipez))
- cleanup objects [\#22](https://github.com/Flipez/rocket-lang/pull/22) ([Flipez](https://github.com/Flipez))
- Boost objects [\#20](https://github.com/Flipez/rocket-lang/pull/20) ([Flipez](https://github.com/Flipez))
- Improve coverage [\#19](https://github.com/Flipez/rocket-lang/pull/19) ([Flipez](https://github.com/Flipez))
- Omit null [\#23](https://github.com/Flipez/rocket-lang/pull/23) ([Flipez](https://github.com/Flipez))
- Remove builtins [\#21](https://github.com/Flipez/rocket-lang/pull/21) ([Flipez](https://github.com/Flipez))
- Foreach [\#18](https://github.com/Flipez/rocket-lang/pull/18) ([Flipez](https://github.com/Flipez))
- Objectmethod arg validation [\#17](https://github.com/Flipez/rocket-lang/pull/17) ([MarkusFreitag](https://github.com/MarkusFreitag))
- Rework methods [\#16](https://github.com/Flipez/rocket-lang/pull/16) ([Flipez](https://github.com/Flipez))
- Release RocketLang 0.10.0 [\#15](https://github.com/Flipez/rocket-lang/pull/15) ([Flipez](https://github.com/Flipez))
- Add object calls [\#14](https://github.com/Flipez/rocket-lang/pull/14) ([Flipez](https://github.com/Flipez))

## [v0.9.7](https://github.com/flipez/rocket-lang/tree/v0.9.7) (2021-09-29)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.9.6...v0.9.7)

**Implemented enhancements:**

- Add `exit()` builtin [\#11](https://github.com/Flipez/rocket-lang/issues/11)
- Add support for `//` as single line comments [\#7](https://github.com/Flipez/rocket-lang/issues/7)
- Add `raise()` builtin [\#5](https://github.com/Flipez/rocket-lang/issues/5)

**Closed issues:**

- Add code coverage [\#8](https://github.com/Flipez/rocket-lang/issues/8)
- Add index operator for strings [\#3](https://github.com/Flipez/rocket-lang/issues/3)

## [v0.9.6](https://github.com/flipez/rocket-lang/tree/v0.9.6) (2021-09-28)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.9.6-rc2...v0.9.6)

**Merged pull requests:**

- Add the first real working script [\#2](https://github.com/Flipez/rocket-lang/pull/2) ([Kjarrigan](https://github.com/Kjarrigan))

## [v0.9.6-rc2](https://github.com/flipez/rocket-lang/tree/v0.9.6-rc2) (2021-09-28)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.9.6-rc1...v0.9.6-rc2)

## [v0.9.6-rc1](https://github.com/flipez/rocket-lang/tree/v0.9.6-rc1) (2021-09-28)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.9.5...v0.9.6-rc1)

## [v0.9.5](https://github.com/flipez/rocket-lang/tree/v0.9.5) (2021-09-27)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/f4fd422c7807b8b00917f983d8399772b7007426...v0.9.5)



\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*
