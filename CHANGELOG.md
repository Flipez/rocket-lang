# Changelog

## [v0.23.0](https://github.com/flipez/rocket-lang/tree/v0.23.0) (2026-01-17)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.22.1...v0.23.0)

**Implemented enhancements:**

- Consider dropping braces requirement for control expressions [\#203](https://github.com/Flipez/rocket-lang/issues/203)
- Support having multiple return values [\#112](https://github.com/Flipez/rocket-lang/issues/112)
- Implement matrix multiplication [\#64](https://github.com/Flipez/rocket-lang/issues/64)
- Add support for `array.each()` [\#44](https://github.com/Flipez/rocket-lang/issues/44)

**Fixed bugs:**

- unique behavior of puts [\#170](https://github.com/Flipez/rocket-lang/issues/170)
- panic: runtime error: makeslice: len out of range [\#166](https://github.com/Flipez/rocket-lang/issues/166)

**Merged pull requests:**

- fix\(int/string\): prevent int downcast [\#229](https://github.com/Flipez/rocket-lang/pull/229) ([Flipez](https://github.com/Flipez))
- fix\(deps\): update dependencies [\#228](https://github.com/Flipez/rocket-lang/pull/228) ([Flipez](https://github.com/Flipez))
- feat\(assign\): add ability to unpack variables on assignment [\#227](https://github.com/Flipez/rocket-lang/pull/227) ([Flipez](https://github.com/Flipez))
- feat\(control\_flow\): make parentheses optional [\#226](https://github.com/Flipez/rocket-lang/pull/226) ([Flipez](https://github.com/Flipez))
- feat\(puts\): improve behavior of puts [\#225](https://github.com/Flipez/rocket-lang/pull/225) ([Flipez](https://github.com/Flipez))
- fix\(array\): fix go panic on .pop\(\) on empty array [\#224](https://github.com/Flipez/rocket-lang/pull/224) ([Flipez](https://github.com/Flipez))
- feat\(docs\):update docusaurus [\#223](https://github.com/Flipez/rocket-lang/pull/223) ([Flipez](https://github.com/Flipez))
- feat\(matrix\) add matrix object [\#222](https://github.com/Flipez/rocket-lang/pull/222) ([Flipez](https://github.com/Flipez))
- chore\(docs\): Update Docusaurus to 3.8 [\#217](https://github.com/Flipez/rocket-lang/pull/217) ([Flipez](https://github.com/Flipez))
- Bump on-headers and compression in /docs [\#215](https://github.com/Flipez/rocket-lang/pull/215) ([dependabot[bot]](https://github.com/apps/dependabot))
- Bump image-size from 1.2.0 to 1.2.1 in /docs [\#212](https://github.com/Flipez/rocket-lang/pull/212) ([dependabot[bot]](https://github.com/apps/dependabot))

## [v0.22.1](https://github.com/flipez/rocket-lang/tree/v0.22.1) (2025-03-15)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.22.0...v0.22.1)

**Implemented enhancements:**

- Implement elseif [\#198](https://github.com/Flipez/rocket-lang/pull/198) ([MarkusFreitag](https://github.com/MarkusFreitag))

**Merged pull requests:**

- fix\(release\): Update goreleaser config to v2 [\#210](https://github.com/Flipez/rocket-lang/pull/210) ([Flipez](https://github.com/Flipez))
- chore\(docs\): Update to Docusaurus 3.7 and React 19 [\#209](https://github.com/Flipez/rocket-lang/pull/209) ([Flipez](https://github.com/Flipez))
- chore\(go\): Update Go to 1.24 [\#208](https://github.com/Flipez/rocket-lang/pull/208) ([Flipez](https://github.com/Flipez))
- Bump express from 4.18.1 to 4.19.2 in /docs [\#207](https://github.com/Flipez/rocket-lang/pull/207) ([dependabot[bot]](https://github.com/apps/dependabot))
- Bump ua-parser-js from 0.7.31 to 0.7.37 in /docs [\#202](https://github.com/Flipez/rocket-lang/pull/202) ([dependabot[bot]](https://github.com/apps/dependabot))
- Bump @babel/traverse from 7.18.10 to 7.23.4 in /docs [\#201](https://github.com/Flipez/rocket-lang/pull/201) ([dependabot[bot]](https://github.com/apps/dependabot))
- feat\(docusaurus\): Update docusaurus to 3.0.0 [\#199](https://github.com/Flipez/rocket-lang/pull/199) ([Flipez](https://github.com/Flipez))
- Bump json5 from 2.2.1 to 2.2.3 in /docs [\#182](https://github.com/Flipez/rocket-lang/pull/182) ([dependabot[bot]](https://github.com/apps/dependabot))

## [v0.22.0](https://github.com/flipez/rocket-lang/tree/v0.22.0) (2023-09-30)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.21.0...v0.22.0)

**Implemented enhancements:**

- \[object/array\] Rename `.yeet()` and `.yoink()` to `.pop()` and `.push()` [\#179](https://github.com/Flipez/rocket-lang/pull/179) ([Flipez](https://github.com/Flipez))
- \[object/array\] Add ability to `.sum()` elements [\#178](https://github.com/Flipez/rocket-lang/pull/178) ([Flipez](https://github.com/Flipez))
- \[object/array\] Add ability to `.join()` elements [\#177](https://github.com/Flipez/rocket-lang/pull/177) ([Flipez](https://github.com/Flipez))
- \[object\] Refactor type conversion and rename methods [\#175](https://github.com/Flipez/rocket-lang/pull/175) ([Flipez](https://github.com/Flipez))
- \[language/control-expressions\] Add support for `🚀-range`  syntax [\#174](https://github.com/Flipez/rocket-lang/pull/174) ([Flipez](https://github.com/Flipez))
- \[errorhandling\] Fix line position, add `file:line:pos` to multiple error messages [\#173](https://github.com/Flipez/rocket-lang/pull/173) ([Flipez](https://github.com/Flipez))

**Fixed bugs:**

- \[evaluator/while\]: Add support for `break` [\#168](https://github.com/Flipez/rocket-lang/pull/168) ([MarkusFreitag](https://github.com/MarkusFreitag))

**Merged pull requests:**

- Bump go version to 1.21 [\#195](https://github.com/Flipez/rocket-lang/pull/195) ([Flipez](https://github.com/Flipez))

## [v0.21.0](https://github.com/flipez/rocket-lang/tree/v0.21.0) (2022-12-06)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.20.1...v0.21.0)

**Implemented enhancements:**

- \[object\] Add `.include?` to ARRAY and HASH [\#165](https://github.com/Flipez/rocket-lang/pull/165) ([MarkusFreitag](https://github.com/MarkusFreitag))
- \[object/string\]: add ascii conversion method [\#164](https://github.com/Flipez/rocket-lang/pull/164) ([MarkusFreitag](https://github.com/MarkusFreitag))
- \[object/array\]: add slices method [\#163](https://github.com/Flipez/rocket-lang/pull/163) ([MarkusFreitag](https://github.com/MarkusFreitag))
- \[object/hash\]: add get method with default if key not exists [\#162](https://github.com/Flipez/rocket-lang/pull/162) ([MarkusFreitag](https://github.com/MarkusFreitag))

## [v0.20.1](https://github.com/flipez/rocket-lang/tree/v0.20.1) (2022-11-01)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.20.0...v0.20.1)

**Implemented enhancements:**

- \[object/string\] Add support for HASH and ARRAY to `.format()` [\#148](https://github.com/Flipez/rocket-lang/pull/148) ([MarkusFreitag](https://github.com/MarkusFreitag))

**Fixed bugs:**

- \[object/string\] Fix `.find()` and `.count()`  argument validation to only accept STRING [\#147](https://github.com/Flipez/rocket-lang/pull/147) ([Flipez](https://github.com/Flipez))

## [v0.20.0](https://github.com/flipez/rocket-lang/tree/v0.20.0) (2022-11-01)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.19.1...v0.20.0)

**Implemented enhancements:**

- \[object/error\] Add `raise()` builtin [\#145](https://github.com/Flipez/rocket-lang/pull/145) ([Flipez](https://github.com/Flipez))
- \[object/error\] Add ability to rescue errors and introduce `begin/rescue/end` [\#142](https://github.com/Flipez/rocket-lang/pull/142) ([Flipez](https://github.com/Flipez))

**Fixed bugs:**

- \[object/funtion\] Missing name to use when inspect [\#143](https://github.com/Flipez/rocket-lang/pull/143) ([Flipez](https://github.com/Flipez))

## [v0.19.1](https://github.com/flipez/rocket-lang/tree/v0.19.1) (2022-10-31)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.19.0...v0.19.1)

**Implemented enhancements:**

- \[stdlib/time\] Add support for `Time.format()` and `Time.parse()` [\#140](https://github.com/Flipez/rocket-lang/pull/140) ([Flipez](https://github.com/Flipez))
- \[object/string\] Add support for `.format()` [\#139](https://github.com/Flipez/rocket-lang/pull/139) ([Flipez](https://github.com/Flipez))
- \[object/array\] Add ability to `.reverse()` [\#138](https://github.com/Flipez/rocket-lang/pull/138) ([Flipez](https://github.com/Flipez))
- \[object/array\] Add ability to `.sort()` [\#137](https://github.com/Flipez/rocket-lang/pull/137) ([Flipez](https://github.com/Flipez))

## [v0.19.0](https://github.com/flipez/rocket-lang/tree/v0.19.0) (2022-10-30)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.18.0...v0.19.0)

**Implemented enhancements:**

- \[foreach/hash\] Swap key/value order [\#132](https://github.com/Flipez/rocket-lang/pull/132) ([Flipez](https://github.com/Flipez))
- \[import\] Allow module name to be defined optionally [\#123](https://github.com/Flipez/rocket-lang/pull/123) ([Flipez](https://github.com/Flipez))
- \[stdlib/time\] Add Time standard library [\#113](https://github.com/Flipez/rocket-lang/pull/113) ([Flipez](https://github.com/Flipez))
- \[stdlib/os\] Add standard library OS with exit\(\) and raise\(\) [\#111](https://github.com/Flipez/rocket-lang/pull/111) ([Flipez](https://github.com/Flipez))
- \[stdlib/io\] Add standard library IO with open\(\) [\#109](https://github.com/Flipez/rocket-lang/pull/109) ([Flipez](https://github.com/Flipez))
- \[builtin\] Add Math standard library and rewrite builtins [\#108](https://github.com/Flipez/rocket-lang/pull/108) ([Flipez](https://github.com/Flipez))
- \[language\] Implement `and`, `or` , `&&` and `||` [\#102](https://github.com/Flipez/rocket-lang/pull/102) ([Flipez](https://github.com/Flipez))
- \[docs\] Add Playground to website [\#100](https://github.com/Flipez/rocket-lang/pull/100) ([Flipez](https://github.com/Flipez))

**Fixed bugs:**

- \[evaluator/assign\] Fix assign if assigned to nested element [\#129](https://github.com/Flipez/rocket-lang/pull/129) ([Flipez](https://github.com/Flipez))
- \[foreach\] Add internal object iterator to fix nested loops [\#122](https://github.com/Flipez/rocket-lang/pull/122) ([MarkusFreitag](https://github.com/MarkusFreitag))
- \[lexer,parser\] Fix line calculation in error messages [\#117](https://github.com/Flipez/rocket-lang/pull/117) ([Flipez](https://github.com/Flipez))
- \[repl\] Fix bugged command history in repl [\#107](https://github.com/Flipez/rocket-lang/pull/107) ([Flipez](https://github.com/Flipez))

**Merged pull requests:**

- \[docs\] Migrate Website to Docusaurus [\#97](https://github.com/Flipez/rocket-lang/pull/97) ([Flipez](https://github.com/Flipez))

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

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.16.0-alpha.3...v0.16.0)

**Merged pull requests:**

- Add support for `next()` and `break()` [\#90](https://github.com/Flipez/rocket-lang/pull/90) ([Flipez](https://github.com/Flipez))

## [v0.16.0-alpha.3](https://github.com/flipez/rocket-lang/tree/v0.16.0-alpha.3) (2022-06-27)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.16.0-alpha.2...v0.16.0-alpha.3)

**Implemented enhancements:**

- Improve file.write\(\) [\#37](https://github.com/Flipez/rocket-lang/issues/37)

**Merged pull requests:**

- Remove support for curly braces in foreach, while, if and function [\#89](https://github.com/Flipez/rocket-lang/pull/89) ([Flipez](https://github.com/Flipez))
- Return bytesWritten instead of True on successfull file.write\(\) [\#88](https://github.com/Flipez/rocket-lang/pull/88) ([Flipez](https://github.com/Flipez))

## [v0.16.0-alpha.2](https://github.com/flipez/rocket-lang/tree/v0.16.0-alpha.2) (2022-04-26)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.16.0-alpha.1...v0.16.0-alpha.2)

**Implemented enhancements:**

- add basic http response [\#84](https://github.com/Flipez/rocket-lang/pull/84) ([Flipez](https://github.com/Flipez))

**Merged pull requests:**

- Apt repo [\#87](https://github.com/Flipez/rocket-lang/pull/87) ([Flipez](https://github.com/Flipez))

## [v0.16.0-alpha.1](https://github.com/flipez/rocket-lang/tree/v0.16.0-alpha.1) (2022-03-14)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.16.0-alpha...v0.16.0-alpha.1)

**Implemented enhancements:**

- add .to\_json\(\) [\#82](https://github.com/Flipez/rocket-lang/pull/82) ([Flipez](https://github.com/Flipez))

## [v0.16.0-alpha](https://github.com/flipez/rocket-lang/tree/v0.16.0-alpha) (2022-03-13)

[Full Changelog](https://github.com/flipez/rocket-lang/compare/v0.15.1...v0.16.0-alpha)

**Implemented enhancements:**

- Fix indexing [\#79](https://github.com/Flipez/rocket-lang/issues/79)
- add networking [\#81](https://github.com/Flipez/rocket-lang/pull/81) ([Flipez](https://github.com/Flipez))
- add ast tests [\#77](https://github.com/Flipez/rocket-lang/pull/77) ([Flipez](https://github.com/Flipez))
- Implement while loop [\#75](https://github.com/Flipez/rocket-lang/pull/75) ([MarkusFreitag](https://github.com/MarkusFreitag))
- add ternary [\#73](https://github.com/Flipez/rocket-lang/pull/73) ([Flipez](https://github.com/Flipez))
- object/integer: satisfy iterable interface [\#66](https://github.com/Flipez/rocket-lang/pull/66) ([RaphaelPour](https://github.com/RaphaelPour))
- Improve Infix operator [\#65](https://github.com/Flipez/rocket-lang/pull/65) ([Flipez](https://github.com/Flipez))

**Fixed bugs:**

- Fix indexable object types [\#80](https://github.com/Flipez/rocket-lang/pull/80) ([Flipez](https://github.com/Flipez))

**Merged pull requests:**

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
