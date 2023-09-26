ChangeLog
---------

English version [by Google](https://gitflic-ru.translate.goog/project/erthink/libmdbx/blob?file=ChangeLog.md&_x_tr_sl=ru&_x_tr_tl=en)
and [by Yandex](https://translated.turbopages.org/proxy_u/ru-en.en/https/gitflic.ru/project/erthink/libmdbx/blob?file=ChangeLog.md).

## v0.12.7 "artec" of 2023-06-16

Stabilizing release with correction of identified errors and shortcomings on the day of the founding of the International Children's Centre[Artek](https://ru.wikipedia.org/wiki/%D0%90%D1%80%D1%82%D0%B5%D0%BA "").

    14 files changed, 222 insertions(+), 56 deletions(-)
    Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
    </leo@yuriev.ru>

Corrections and refinements:

- Fixing a misprint in the variable's name inside `mdbx_env_turn_for_recovery()` which led to misbehaviour in certain situations. From a users point of view, given the current Utilities scenarios `mdbx_chk` there was only one specific/received error/issue scenario - when the weak/weak meta page was checked and activated with НЕ-последней transactions after the car's system crash where the OBD was used in fragile/unsafe mode. In the scenario, when the target page was successfully checked and activated, an error was reported relating to the non-coherence control mechanism of the file system and the OBD data displayed in the DOS. The OBD was successfully restored and there were no negative consequences except the error report itself. Technically, the error occurred when the metapage was "shifted" when one of the other two metapages had a larger transaction number:
- While the content of other meta-pages was correct and the transaction numbers were larger, the resulting transaction number in the target/activated meta-page was not appropriate to these meta-pages and could be less or equal.
- As a result, if such meta-pages were weak/weak status, the anti-coherence protection unified buffer/page cache could be activated when the OBD was closed after switching, and an assert check could be activated in debugging assemblies.
- If such meta-pages were strong/style, the switch to a new meta-page might not have had an effect or resulted in two meta-pages with the same transaction number, which is a wrong situation.
- Overriding assembly problems through GCC using options `-m32 -arch=i686 -Ofast`.The problem is caused by the GCC error that caused the design `__attribute__((__target__("sse2")))` does not include full use of the SSE and SSE2 instructions unless this was done by command line options, but the option was used `-Ofast`.As a result, the assembly ended with an error report:`error: inlining failed in call to 'always_inline' '_mm_movemask_ps': target specific option mismatch`
- Finalization of the OBD "rehabilitation" mode and switching to the specified meta-page:
   - Elimination of the update without the need for a meta-page with an increase in transaction number;
   - Removal of a pointless/excessive warning from the OBD geometry update;
   - More expected and safe behaviour when checking OBD with target meta-page in reading-record mode.
- Now when the OBD is opened by `mdbx_env_open_for_recovery()` The OBD shall not be tampered with, including when the OBD is closed. This shall make it possible to secure the OBD (destroy the probability of its destruction) if the user, in an attempt to recover, or simply as an experiment, has delivered a disposal device. `mdbx_chk` A combination of incorrect or dangerous parameters is still in operation, and the normal verification, like the explicit switching of meta-pages, is still in operation.

Little things:

- Minor specification CMake-пробника for `std::filesystem`, checking the need for a lincoque with additional libraries C++.
- Elimination of minor warnings from old compilers in tests.
- Addressing the cause of false-positive warnings of new versions of GCC in C++API.
- Correction of the link to the benchmarking repository of the ioarena.
- Add cross-references to the doxygen documentation for C++ API.
- Clarification of restrictions in section[Restructions &amp; Caveats](https://libmdbx.dqdkfa.ru/intro.html#restrictions "").
- Correction of references to description `mdbx_canary_put()`.

------------------------------------------------------------------------------------------------------------------------------

## v0.12.6 "Scene" of 2023-04-29

Stabilizing issue with correction of identified errors and elimination of deficiencies, on the day of the 100-year anniversary of the sports club[«ЦСКА»](https://ru.wikipedia.org/wiki/%D0%A6%D0%B5%D0%BD%D1%82%D1%80%D0%B0%D0%BB%D1%8C%D0%BD%D1%8B%D0%B9_%D1%81%D0%BF%D0%BE%D1%80%D1%82%D0%B8%D0%B2%D0%BD%D1%8B%D0%B9_%D0%BA%D0%BB%D1%83%D0%B1_%D0%90%D1%80%D0%BC%D0%B8%D0%B8 "").

    14 files changed, 117 insertions(+), 83 deletions(-)
    Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
    </leo@yuriev.ru>

Little things:

- Update the patch for old version of the bildroot.
- Use of clang-format-16.
- Use `enum` - Types instead of `int` to eliminate GCC 13 warnings, which could break the assembly in Fedora 38.

------------------------------------------------------------------------------------------------------------------------------

## v0.12.5 "Dynam" of 2023-04-18

Stabilizing issue with correction of identified errors and shortcomings, 100-year anniversary day of sports society[Dynamo](https://ru.wikipedia.org/wiki/%D0%94%D0%B8%D0%BD%D0%B0%D0%BC%D0%BE_(%D1%81%D0%BF%D0%BE%D1%80%D1%82%D0%B8%D0%B2%D0%BD%D0%BE%D0%B5_%D0%BE%D0%B1%D1%89%D0%B5%D1%81%D1%82%D0%B2%D0%BE) "").

    16 files changed, 686 insertions(+), 247 deletions(-)
    Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
    </leo@yuriev.ru>

Thanks:

- Max[maxc0d3r@protonmail.com](mailto:maxc0d3r@protonmail.com "")For reporting on the issue of exports from DSO/DLL obsolete functions API.
- `@calvin3721` for reporting on work-related issues `MainDB` With flags, it's not quiet.

Corrections:

- Export from DSO/DLL is corrected for obsolete functions that are replaced online in the current API.
- The use of an incorrect comparator in the creation or re-establishment has been eliminated `MainDB` with flags/opposions involving a specific comparator (not silent).

Little things:

- Duplicate diagnostics inside removed `node_read_bigdata()`.
- References in the description corrected `mdbx_env_set_geometry()`.
- A separate test added `extra/upsert_alldups` For a specific replacement/rewrite scenario, a single value of all multi-purposes of the relevant key, i.e., the replacement of all "duplicates" with one value.
- Options added to C++ API `buffer::key_from()` with a clear name by type of data.
- A separate test added `extra/maindb_ordinal` for a specific creation scenario `MainDB` with flags requiring the use of a comparator in a non-silence manner.
- Reactoring the "coherence" test of meta pages.
- Adjustment `osal_vasprintf()` to eliminate static analyzer warnings.

------------------------------------------------------------------------------------------------------------------------------

## v0.12.4 "art. 333" of 2023-03-03

Stabilizing output with corrected errors, corrected deficiencies and technical debts: Line 0.12 is considered ready for use, is stable and will only receive error correction thereafter. Development will continue in branch 0.13 and branch 0.11 becomes archival.

    63 files changed, 1161 insertions(+), 569 deletions(-)
    Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
    </leo@yuriev.ru>

Thanks:

- Max[maxc0d3r@protonmail.com](mailto:maxc0d3r@protonmail.com "")for the message on the problem ERROR_SHARING_VIOLATION on the MDBX_EXCLUSIVE on Windows.
- Alisher Ashyrov[https://t.me/a1is43ras4](https://t.me/a1is43ras4 "")for reporting the problem with the assert check and assisting in debugging.
- Masatoshi Fukunaga[https://gitflic.ru/user/mah0x211](https://gitflic.ru/user/mah0x211 "")for reporting on the problem`put(MDBX_UPSERT+MDBX_ALLDUPS)`For the case of replacement of all values of subDb.

Corrections:

- The regression following the 474391c83c5f81def6fdf3b0b6f5716a87b78bfff has been eliminated, resulting in the return of ERROR_SARING_VIOLATION to Windows when the OBD is opened in the MDBX_EXLUSIVE mode for reading-recording.
- Added a display size limit at a short read-only file to prevent error ERROR_NOT_ENOUGH_MEMORY in Windows, which is then not informative at all for the user.
- Reactorization conducted `dxb_resize()` including, to eliminate the response of the ASERT check `size_bytes == env-&gt;me_dxb_mmap.current` The check only worked in debugging assemblies, with specific reading and writing transactions in different streams, at the same time as the OBD size was changed. In addition to operating the check, there were no other consequences.
- The problem with`put(MDBX_UPSERT+MDBX_ALLDUPS)` In this operation, subDb becomes completely empty, without any pages, and this is the situation that has not been included in the code, which has resulted in OBD damage when the transaction is fixed.
- Over-assert check inside removed `override_meta()`.Which in debugging assemblies could lead to false responses to OBD recovery, including automatic rebound of weak meta-pages.
- Macroeconomics adjusted `__cold` / `__hot` including to address the problem `error: inlining failed in call to ‘always_inline FOO(...)’: target specific option mismatch` when assembled using GCC &gt;10.x for SH4.


Elimination of technical debts and minorities:

- Numerous errors in the documentation have been corrected.
- Completed test for full stochastic testing `MDBX_EKEYMISMATCH` in mode `MDBX_APPEND`.
- Launch scenarios expanded `mdbx_chk` CMake-тестов for verification in both normal and exclusive reading and recording modes.
- Refined Specification `const` and `noexcept` For several methods in C++ API.
- The use of the drain under the buffers has been eliminated `wchar` - route transformation.
- For Windows, a function added `mdbx_env_get_path()` to obtain a path to the OBD in multibyte character format.
- Added doxygen descriptions for API with broad symbols.
- The MSVC static analyzer warnings have been removed, all of which were irrelevant or false.
- The false GCC warning for SH4 was removed.
- The ASAN (Address Sanitizer) support for the MSVC assembly has been added.
- Script resample set expanded `test/long_stochastic.sh` the option added `--extra`.
- C++ API Added support for extended implementation time options `mdbx::extra_runtime_option` similar `enum MDBX_option_t` C API.
- Discharge of all page-operations counters `mdbx_stat`.

------------------------------------------------------------------------------------------------------------------------------

## v0.12.3 "akul" of 2023-01-07

Output with significant refinements and new functionality to remember closed open-source[The Shark Project](https://erigon.substack.com/p/winding-down-support-for-akula-project "").

Added prefault recording, reset control of uncoherence. unified page/buffer cache, changed the way the pages were merged, etc.

    20 files changed, 4508 insertions(+), 2928 deletions(-)
    Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
    </leo@yuriev.ru>

Thanks:

- [Alex Sharov](https://t.me/AskAlexSharov "")and team[Erigon](https://github.com/ledgerwatch/erigon "")for testing.
- [Simon Leier](https://t.me/leisim "")for reporting malfunctions and testing.

New:

- Use of address[https://libmdbx.dqdkfa.ru/dead](https://libmdbx.dqdkfa.ru/dead "")-github to refer to stored in web.archive.org copies of resources destroyed by the Github administration.

- Prefault recording for read-write display pages has been implemented. This results in multiple reductions in system costs and significant productivity gains in the respective use scenarios when:
   - The size of the OBD and data volume is significantly higher than the LSA;
   - mode used `MDBX_WRITEMAP`;
   - Non-small transactions (in progress, many hundreds or thousands of pages are identified).
- In mode `MDBX_WRITEMAP` The selection/re-use of pages results in page-fault and reading from the disk, even if the content of the page is not necessary (to be rewritten). This is the result of the virtual memory subsystem and the regular mode of treatment through `MADV_REMOVE` It does not work for all FSs and is usually more expensive than savings.
   - Now libmdbx uses a "precautionary entry" of pages that are on systems with[Unified page cache](https://www.opennet.ru/base/dev/ubc.txt.html "")results in data "pulling" by eliminating the need to read from the disk when accessing such a memory page.
   - The new functionality works in harmony with automatic read-ahed control and cache of the presence status of pages in the OZM, through[mincore()](https://man7.org/linux/man-pages/man2/mincore.2.html "").
- An option added `MDBX_opt_prefault_write_enable` For the possibility of mandatory inclusion/deactivation of prefault recording.
- Dynamic choice made between the disk and the regular recording followed by[fdatasync()](https://man7.org/linux/man-pages/man3/fdatasync.3p.html "")managed by an option `MDBX_opt_writethrough_threshold`. In long-term (durable) modes, data can be released on a disc in two ways:
   - Through the file descriptor open with `O_DSYNC`;
   - normal recording followed by call `fdatasync()`.
- The first option is better for writing down a small number of pages and/or if the channel with the disc/host has close to zero delay. The second option is better if many pages are required and/or the channel has a significant delay (date centres, clouds). `MDBX_opt_writethrough_threshold` Allows a threshold to be set at the time of execution for the dynamic choice of the mode of recording, depending on the volume and the specific conditions of use.
- Automatic installation `MDBX_opt_rp_augment_limit` Depending on the size of the OBD.
- Prohibition of different treatment `MDBX_WRITEMAP` between processes in deferred/lazy recording modes, because in this case it is not possible to release data on a disc in all cases on all supported platforms.
- Assembly option added `MDBX_MMAP_USE_MS_ASYNC` Disabled system call `msync(MS_ASYNC)`which is not necessary on the vast majority of relevant LOs. `MDBX_MMAP_USE_MS_ASYNC=0`(deleted) on Linux and other systems with unified page cache. `msync(MS_ASYNC)`) corresponds to the unaltered logic of LMDB. As a result, in simple/naive-based benchmarking, libmdbx exceeds LMDB approximately as much as in real use. Just in case, it should be noted/remember that on Windows, it is assumed that libmdbx will lag behind LMDB in multiple small transaction scenarios, because libmdbx knowingly uses file locks on Windows that are slow (badly implemented in the OS core) but allow users to be insured against a mass of incorrect actions causing OBD damage.
- Support for non-print names for subDb.
- A clear choice added `tls_model("local-dynamic")` to avoid the problem. `relocation R_X86_64_TPOFF32 against FOO cannot be used with -shared`due to an error in the CLANG leading to the use of an incorrect mode `ls_model`.
- Change in the method of merging pages at removal. The merger is now performed mainly with the already modified/filtered page. If both pages on the right and on the left are of the same status, then the least completed page is the same as before. In mass disposal scenarios, this increases productivity to 50 per cent.
- Added LCK-file absence check with alternative name.

Corrections (without adjustment of new functions):

- Change of display size if required to release data on disk when called`mdbx_env_sync()`from a parallel flow of performance outside the operating transaction.
- Correction of the regression after the comedine db72763de049d6e4546f838277fe83b9081ad1de of 2022-10-08 in the logic of returning dirty pages in mode`MDBX_WRITEMAP`that led to the use of the free pages not immediately, but rather to a retried transaction list and an unjustified increase in transaction size.

- Removal of SIGSGV or erroneous call `free()` in situations of re-opening by `mdbx_env_open()`.
- Remedial of an error made in the comms fe20de136c22ed3bc4c6d3f673e79c106e824f60 of 2022-09-18, resulting in Linux mode `MDBX_WRITEMAP` Never called. `msync()`. The problem exists only in release 0.12.2.
- Adding dirty pages to `MDBX_WRITEMAP` to be made available by `mdbx_txn_info()` up-to-date information on the extent of changes in reading-recording transactions.
- Correction of non-existent typo under `#if` The order of byte.
- Correction of assembly for incident `MDBX_PNL_ASCENDING=1`.


Elimination of technical debts and minorities:

- Further development of support for the internal integration of GC records`page_alloc_slowpath()`.
- Removal of minor warnings of Convergence.
- Use of a single locator to search GC.
- Reprocessing of internal flags related to GC pages.
- Finalize the preparation of the reserve before updating GC with BigFoot included.
- Optimization `pnl_merge()` For non-overlapping combined lists.
- Optimizing support to the graded page list `dpl_append()`.
- Accelerating work `mdbx_chk` when processing user records in `@MAIN`.
- Reprocessing LRU spilling tags.
- Re-engineering of the control of "uncoherence" Unified page cache to reduce overhead costs.
- Reactorization and microoptimization.

------------------------------------------------------------------------------------------------------------------------------

## v0.12.2 "Ivan Yargin" of 2022-11-11

Output with significant refinements and new functionality to remember Russian wrestler[Ivane Sergeyeviche Yarriginé](https://ru.wikipedia.org/wiki/%D0%AF%D1%80%D1%8B%D0%B3%D0%B8%D0%BD,_%D0%98%D0%B2%D0%B0%D0%BD_%D0%A1%D0%B5%D1%80%D0%B3%D0%B5%D0%B5%D0%B2%D0%B8%D1%87 "").

At the Olympic Games in Munich in 1972, Ivan Yarrigin put all the rivals on the shoulder, spending less than nine minutes, a record that has not been broken by anyone until now.

    64 files changed, 5573 insertions(+), 2510 deletions(-)
    Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
    </leo@yuriev.ru>

New:

- Support all major assembly options through CMake.
- Requirements for CMake are lowered to version 3.0.2 to allow assembly for obsolete platforms.
- The possibility of profiling GC work in complex and/or loaded scenarios (e.g. Ethereum/Erigon) has been added. The code has been shut down and the assembly option needs to be specified for activation.`MDBX_ENABLE_PROFGC=1`.
- Added Function `mdbx_env_warmup()` for OBD "heating" with the possibility of keeping the pages in memory.`mdbx_chk`, `mdbx_copy` and `mdbx_dump` Options added `-u` and `-U` to activate the relevant functional unit.
- Disabled treatment of "foul" pages in non-negative modes (in thousands of United States dollars) `MDBX_WRITEMAP` to `MDBX_AVOID_MSYNC=0` Further development has reduced overhead costs and has been planned for a long time, but has been postponed as it has required other changes.

- Removing (spilling) dirty pages with large/overflow pages. Finalization allows correct compliance with options policy `MDBX_opt_txn_dp_limit`,
  `MDBX_opt_spill_max_denominator`, `MDBX_opt_spill_min_denominator`and was planned long ago, but delayed as it required other changes.
- For Windows, UNICODE definitions of macros added in API`MDBX_DATANAME`, `MDBX_LOCKNAME`and `MDBX_LOCK_SUFFIX`.
- Switch to type-use priority `size_t` In order to reduce overhead costs on the Elbrus platform.
- Functions added to API `mdbx_limits_valsize4page_max()` and `mdbx_env_get_valsize4page_max()`Returning the maximum size in bytes, which may be placed in the same large/overflow page, rather than in sequences of two or more such pages. For tables supporting duplicates, taking out values on the large/overflow page is not supported, so the result is consistent with`mdbx_limits_valsize_max()`.
- Functions added to API `mdbx_limits_pairsize4page_max()` and `mdbx_env_get_pairsize4page_max()` Returning the maximum total size of the pair of key-values in bytes to be placed on the same page, without removing the value to a separate large/overflow page. For tables supporting duplicates, the large/overflow page is not supported, so the result determines the maximum/acceptable total size of the key-value pair.
- Use of asynchronous (overlapped) Windows entry, including non-bufied input and output `WriteGather()`.This reduces overhead costs and partially circumvents Windows problems with low input-output productivity, including high delays`FlushFileBuffers()`.The new code also consolidates the recorded regions on all platforms, while on Windows, the use of events (events) is reduced to a minimum while automatically using`WriteGather()`.Therefore, there is a significant reduction in overhead costs for OS, and in Windows this acceleration, in some scenarios, may be multiple compared to LMDB.
- Assembly option added `MDBX_AVOID_MSYNC` which determines the behavior of libmdbx in mode `MDBX_WRITE_MAP` (when data are modified directly on the OBD pages displayed in the OSL):
   - If `MDBX_AVOID_MSYNC=0` (by default for all systems except Windows), (as in the past) saves data by`msync()`either`FlushViewOfFile()`On Windows platforms with a full virtual memory subsystem and an adequate file output input, this provides a minimum overhead (one system challenge) and maximum productivity. However, Windows is causing significant degradation, including because of post-system-based data-processing problems.`FlushViewOfFile()`Also required to challenge`FlushFileBuffers()`With a mass of problems and fuss inside the OS core.
   - If `MDBX_AVOID_MSYNC=1` (by default only on Windows), data retention is performed by an explicit entry into each amended OBD page file. This requires additional overhead costs, both for tracking amended pages (the maintenance of "Dirty" Page lists) and for system calls for their recording. In addition, in terms of the virtual memory subsystem of the OS kernel, the OBD pages modified in the OSP and clearly recorded in the file, may either remain "Dirty" and be rewritten by the OS kernel later, or require additional overhead costs to track, modify and supplement the data. However, on Windows, this data recording path generally provides higher productivity.
- Improved heuristics for the integration of GC auto-merger.
- Change of LCK format and semantics in some interior fields. Libmdbx versions using different format will not be able to operate with the same OBD at the same time, but only in turn (LCK file rewrites when opening the first OBD opening process).
- `C++` API Added methods of recording transactions with information on delays.
- Added `MDBX_HAVE_BUILT IN_CPU_SUPPORTS` Build option to control use GCC's`__builtin_cpu_supports()`function, which could be unavailable on a fake OSes (macos, ios, android, etc).

Corrections (without adjustment of the above new functions):

- Remove a number of warnings when assembled by MinGW.
- Elimination of false-positive messages from Valgrind about the use of uninitiated data due to levelling gaps in the`struct troika`.
- Corrected return of unexpected error `MDBX_BUSY` from functions `mdbx_env_set_option()`,
  `mdbx_env_set_syncbytes()` and `mdbx_env_set_syncperiod()`.
- Small corrections for compatibility with CMake 3.8
- More control and caution (paranoia) for defects insurance`mremap()`.
- A crutch to repair the old versions `stdatomic.h` from GNU Lib C, where the macros `ATOMIC_*_LOCK_FREE` It's a mistake to re-determine over functions.
- Use `fcntl64(F_GETLK64/F_SETLK64/F_SETLKW64)` This solves the problem of the verification approval response in the assembly of platforms where the type`off_t`wider than the corresponding fields`структуры flock`used to lock files.
- The collection of information on delays in recording transactions has been further developed:
   - Deviation of the measurement of the duration of the update of the GC when the debugging internal audit is included has been eliminated;
   - Protection against undeflow-zero only for total delay in metrics to avoid situations where the sum of individual stages is greater than the total length.
- A number of corrections to eliminate the operation of the verification approval in debugging assemblies.
- More cautious conversion to type `mdbx_tid_t` to eliminate warnings.
- Correction of unnecessary data release to disk mode `MDBX_SAFE_NOSYNC` (b) To update GC.
- Fixed an extra check for `MDBX_APPENDDUP` inside `mdbx_cursor_put()` What could effect in returning `MDBX_EKEYMISMATCH` for valid cases.
- Fixed Nasty `clz()` bug `_BitScanReverse()`, only MSVC is active).

Little things:

- The historical references to a gythub-deleted project are diverted to[Web.archive.org](https://web.archive.org/web/https://github.com/erthink/libmdbx "").
- Synchronized CMake designs between projects.
- A warning of the security of RICS-V has been added.
- Added description of parameters `MDBX_debug_func` and `MDBX_debug_func`.
- A bypass solution has been added to minimize false-positive conflicts when using file blocks in Windows.
- Verification of the atomity of C11 operations with 32/64-bit data.
- 42-fold reduction in silence for `me_options.dp_limit` in the debugging assemblies.
- Adding the platform `gcc-riscv64-linux-gnu` List for purpose `cross-gcc`.
- Small Script Edits `long_stochastic.sh` for Windows.
- Removal of an unnecessary call `LockFileEx()` Inside `mdbx_env_copy()`.
- A description of the use of file descriptors in different modes has been added.
- Use added `_CrtDbgReport()` in the debugging assemblies.
- Fixed an extra ensure/assertion check of`oldest_reader`inside`txn_end()`.
- Remodel definition of deprecated use of`MDBX_NODUPDATA`.
- Fixed decision ASAN/Valgring-under-buildings.
- Fixed minor MingGW Warning.

-------------------------------------------------------------------------------

## v0.12.1 "Positive Proxima" at 2022-08-24

The planned frontward release with new superior features on the day of 20 anniversary of [Positive Technologies](https://ptsecurty.com).

```
37 files changed, 7604 insertions(+), 7417 deletions(-)
Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
```

New:

- The `Big Foot` feature which significantly reduces GC overhead for processing large lists of retired pages from huge transactions.
  Now _libmdbx_ avoid creating large chunks of PNLs (page number lists) which required a long sequences of free pages, aka large/overflow pages.
  Thus avoiding searching, allocating and storing such sequences inside GC.
- Improved hot/online validation and checking of database pages both for more robustness and performance.
- New solid and fast method to latch meta-pages called `Troika`.
  The minimum of memory barriers, reads, comparisons and conditional transitions are used.
- New `MDBX_VALIDATION` environment options to extra validation of DB structure and pages content for carefully/safe handling damaged or untrusted DB.
- Accelerated ×16/×8/×4 by AVX512/AVX2/SSE2/Neon implementations of search page sequences.
- Added the `gcrtime_seconds16dot16` counter to the "Page Operation Statistics" that accumulates time spent for GC searching and reclaiming.
- Copy-with-compactification now clears/zeroes unused gaps inside database pages.
- The `C` and `C++` APIs has been extended and/or refined to simplify using `wchar_t` pathnames.
  On Windows the `mdbx_env_openW()`, `mdbx_env_get_pathW()`, `mdbx_env_copyW()`, `mdbx_env_open_for_recoveryW()` are available for now,
  but the `mdbx_env_get_path()` has been replaced in favor of `mdbx_env_get_pathW()`.
- Added explicit error message for Buildroot's Microblaze toolchain maintainers.
- Added `MDBX_MANAGE_BUILD_FLAGS` build options for CMake.
- Speed-up internal `bsearch`/`lower_bound` implementation using branchless tactic, including workaround for CLANG x86 optimiser bug.
- A lot internal refinement and micro-optimisations.
- Internally counted volume of dirty pages (unused for now but for coming features).

Fixes:

- Never use modern `__cxa_thread_atexit()` on Apple's OSes.
- Don't check owner for finished transactions.
- Fixed typo in `MDBX_EINVAL` which breaks MingGW builds with CLANG.


## v0.12.0 at 2022-06-19

Not a release but preparation for changing feature set and API.


********************************************************************************


## v0.11.14 "Sergey Kapitsa" at 2023-02-14

The stable bugfix release in memory of [Sergey Kapitsa](https://en.wikipedia.org/wiki/Sergey_Kapitsa) on his 95th birthday.

```
22 files changed, 250 insertions(+), 174 deletions(-)
Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
```

Fixes:
- backport: Fixed insignificant typo of `||` inside `#if` byte-order condition.
- backport: Fixed `SIGSEGV` or an erroneous call to `free()` in situations where
  errors occur when reopening by `mdbx_env_open()` of a previously used
  environment.
- backport: Fixed `cursor_put_nochecklen()` internals for case when dupsort'ed named subDb
  contains a single key with multiple values (aka duplicates), which are replaced
  with a single value by put-operation with the `MDBX_UPSERT+MDBX_ALLDUPS` flags.
  In this case, the database becomes completely empty, without any pages.
  However exactly this condition was not considered and thus wasn't handled correctly.
  See [issue#8](https://gitflic.ru/project/erthink/libmdbx/issue/8) for more information.
- backport: Fixed extra assertion inside `override_meta()`, which could
  lead to false-positive failing of the assertion in a debug builds during
  DB recovery and auto-rollback.
- backport: Refined the `__cold`/`__hot` macros to avoid the
  `error: inlining failed in call to ‘always_inline FOO(...)’: target specific option mismatch`
  issue during build using GCC >10.x for SH4 arch.

Minors:

- backport: Using the https://libmdbx.dqdkfa.ru/dead-github
  for resources deleted by the Github' administration.
- backport: Fixed English typos.
- backport: Fixed proto of `__asan_default_options()`.
- backport: Fixed doxygen-description of C++ API, especially of C++20 concepts.
- backport: Refined `const` and `noexcept` for few C++ API methods.
- backport: Fixed copy&paste typo of "Getting started".
- backport: Update MithrilDB status.
- backport: Resolve false-posirive `used uninitialized` warning from GCC >10.x
  while build for SH4 arch.


--------------------------------------------------------------------------------


## v0.11.13 at "Swashplate" 2022-11-10

The stable bugfix release in memory of [Boris Yuryev](https://ru.wikipedia.org/wiki/Юрьев,_Борис_Николаевич) on his 133rd birthday.

```
30 files changed, 405 insertions(+), 136 deletions(-)
Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
```

Fixes:

- Fixed builds with older libc versions after using `fcntl64()` (backport).
- Fixed builds with  older `stdatomic.h` versions,
  where the `ATOMIC_*_LOCK_FREE` macros mistakenly redefined using functions (backport).
- Added workaround for `mremap()` defect to avoid assertion failure (backport).
- Workaround for `encryptfs` bug(s) in the `copy_file_range` implementation  (backport).
- Fixed unexpected `MDBX_BUSY` from `mdbx_env_set_option()`, `mdbx_env_set_syncbytes()`
  and `mdbx_env_set_syncperiod()` (backport).
- CMake requirements lowered to version 3.0.2 (backport).

Minors:

- Minor clarification output of `--help` for `mdbx_test` (backport).
- Added admonition of insecure for RISC-V (backport).
- Stochastic scripts and CMake files synchronized with the `devel` branch.
- Use `--dont-check-ram-size` for small-tests make-targets (backport).


--------------------------------------------------------------------------------


## v0.11.12 "Эребуни" at 2022-10-12

The stable bugfix release.

```
11 files changed, 96 insertions(+), 49 deletions(-)
Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
```

Fixes:

- Fixed static assertion failure on platforms where the `off_t` type is wider
  than corresponding fields of `struct flock` used for file locking (backport).
  Now _libmdbx_ will use `fcntl64(F_GETLK64/F_SETLK64/F_SETLKW64)` if available.
- Fixed assertion check inside `page_retire_ex()` (backport).

Minors:

- Fixed `-Wint-to-pointer-cast` warnings while casting to `mdbx_tid_t` (backport).
- Removed needless `LockFileEx()` inside `mdbx_env_copy()` (backport).


--------------------------------------------------------------------------------


## v0.11.11 "Тендра-1790" at 2022-09-11

The stable bugfix release.

```
10 files changed, 38 insertions(+), 21 deletions(-)
Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
```

Fixes:

- Fixed an extra check for `MDBX_APPENDDUP` inside `mdbx_cursor_put()` which could result in returning `MDBX_EKEYMISMATCH` for valid cases.
- Fixed an extra ensure/assertion check of `oldest_reader` inside `mdbx_txn_end()`.
- Fixed derived C++ builds by removing `MDBX_INTERNAL_FUNC` for `mdbx_w2mb()` and `mdbx_mb2w()`.


--------------------------------------------------------------------------------


## v0.11.10 "the TriColor" at 2022-08-22

The stable bugfix release.

```
14 files changed, 263 insertions(+), 252 deletions(-)
Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
```

New:

- The C++ API has been refined to simplify support for `wchar_t` in path names.
- Added explicit error message for Buildroot's Microblaze toolchain maintainers.

Fixes:

- Never use modern `__cxa_thread_atexit()` on Apple's OSes.
- Use `MultiByteToWideChar(CP_THREAD_ACP)` instead of `mbstowcs()`.
- Don't check owner for finished transactions.
- Fixed typo in `MDBX_EINVAL` which breaks MingGW builds with CLANG.

Minors:

- Fixed variable name typo.
- Using `ldd` to check used dso.
- Added `MDBX_WEAK_IMPORT_ATTRIBUTE` macro.
- Use current transaction geometry for untouched parameters when `env_set_geometry()` called within a write transaction.
- Minor clarified `iov_page()` failure case.


--------------------------------------------------------------------------------


## v0.11.9 "Чирчик-1992" at 2022-08-02

The stable bugfix release.

```
18 files changed, 318 insertions(+), 178 deletions(-)
Signed-off-by: Леонид Юрьев (Leonid Yuriev) <leo@yuriev.ru>
```

Acknowledgments:

- [Alex Sharov](https://github.com/AskAlexSharov) and Erigon team for reporting and testing.
- [Andrew Ashikhmin](https://gitflic.ru/user/yperbasis) for contributing.

New:

- Ability to customise `MDBX_LOCK_SUFFIX`, `MDBX_DATANAME`, `MDBX_LOCKNAME` just by predefine ones during build.
- Added to [`mdbx::env_managed`](https://libmdbx.dqdkfa.ru/group__cxx__api.html#classmdbx_1_1env__managed)'s methods a few overloads with `const char* pathname` parameter (C++ API).

Fixes:

- Fixed hang copy-with-compactification of a corrupted DB
  or in case the volume of output pages is a multiple of `MDBX_ENVCOPY_WRITEBUF`.
- Fixed standalone non-CMake build on MacOS (`#include AvailabilityMacros.h>`).
- Fixed unexpected `MDBX_PAGE_FULL` error in rare cases with large database page sizes.

Minors:

- Minor fixes Doxygen references, comments, descriptions, etc.
- Fixed copy&paste typo inside `meta_checktxnid()`.
- Minor fix `meta_checktxnid()` to avoid assertion in debug mode.
- Minor fix `mdbx_env_set_geometry()` to avoid returning `EINVAL` in particular rare cases.
- Minor refine/fix batch-get testcase for large page size.
- Added `--pagesize NN` option to long-stotastic test script.
- Updated Valgrind-suppressions file for modern GCC.
- Fixed `has no symbols` warning from Apple's ranlib.


--------------------------------------------------------------------------------


## v0.11.8 "Baked Apple" at 2022-06-12

The stable release with an important fixes and workaround for the critical macOS thread-local-storage issue.

Acknowledgments:

- [Masatoshi Fukunaga](https://github.com/mah0x211) for [Lua bindings](https://github.com/mah0x211/lua-libmdbx).

New:

- Added most of transactions flags to the public API.
- Added `MDBX_NOSUCCESS_EMPTY_COMMIT` build option to return non-success result (`MDBX_RESULT_TRUE`) on empty commit.
- Reworked validation and import of DBI-handles into a transaction.
  Assumes  these changes will be invisible to most users, but will cause fewer surprises in complex DBI cases.
- Added ability to open DB in without-LCK (exclusive read-only) mode in case no permissions to create/write LCK-file.

Fixes:

- A series of fixes and improvements for automatically generated documentation (Doxygen).
- Fixed copy&paste bug with could lead to `SIGSEGV` (nullptr dereference) in the exclusive/no-lck mode.
- Fixed minor warnings from modern Apple's CLANG 13.
- Fixed minor warnings from CLANG 14 and in-development CLANG 15.
- Fixed `SIGSEGV` regression in without-LCK (exclusive read-only) mode.
- Fixed `mdbx_check_fs_local()` for CDROM case on Windows.
- Fixed nasty typo of typename which caused false `MDBX_CORRUPTED` error in a rare execution path,
  when the size of the thread-ID type not equal to 8.
- Fixed Elbrus/E2K LCC 1.26 compiler warnings (memory model for atomic operations, etc).
- Fixed write-after-free memory corruption on latest `macOS` during finalization/cleanup of thread(s) that executed read transaction(s).
  > The issue was suddenly discovered by a [CI](https://en.wikipedia.org/wiki/Continuous_integration)
  > after adding an iteration with macOS 11 "Big Sur", and then reproduced on recent release of macOS 12 "Monterey".
  > The issue was never noticed nor reported on macOS 10 "Catalina" nor others.
  > Analysis shown that the problem caused by a change in the behavior of the system library (internals of dyld and pthread)
  > during thread finalization/cleanup: now a memory allocated for a `__thread` variable(s) is released
  > before execution of the registered Thread-Local-Storage destructor(s),
  > thus a TLS-destructor will write-after-free just by legitime dereference any `__thread` variable.
  > This is unexpected crazy-like behavior since the order of resources releasing/destroying
  > is not the reverse of ones acquiring/construction order. Nonetheless such surprise
  > is now workarounded by using atomic compare-and-swap operations on a 64-bit signatures/cookies.

Minors:

- Refined `release-assets` GNU Make target.
- Added logging to `mdbx_fetch_sdb()` to help debugging complex DBI-handels use cases.
- Added explicit error message from probe of no-support for `std::filesystem`.
- Added contributors "score" table by `git fame` to generated docs.
- Added `mdbx_assert_fail()` to public API (mostly for backtracing).
- Now C++20 concepts used/enabled only when `__cpp_lib_concepts >= 202002`.
- Don't provide nor report package information if used as a CMake subproject.


--------------------------------------------------------------------------------


## v0.11.7 "Resurrected Sarmat" at 2022-04-22

The stable risen release after the Github's intentional malicious disaster.

#### We have migrated to a reliable trusted infrastructure
The origin for now is at [GitFlic](https://gitflic.ru/project/erthink/libmdbx)
since on 2022-04-15 the Github administration, without any warning nor
explanation, deleted _libmdbx_ along with a lot of other projects,
simultaneously blocking access for many developers.
For the same reason ~~Github~~ is blacklisted forever.

GitFlic already support Russian and English languages, plan to support more,
including 和 中文. You are welcome!

New:

- Added the `tools-static` make target to build statically linked MDBX tools.
- Support for Microsoft Visual Studio 2022.
- Support build by MinGW' make from command line without CMake.
- Added `mdbx::filesystem` C++ API namespace that corresponds to `std::filesystem` or `std::experimental::filesystem`.
- Created [website](https://libmdbx.dqdkfa.ru/) for online auto-generated documentation.
- Used `https://web.archive.org/web/https://github.com/erthink/libmdbx` for dead (or temporarily lost) resources deleted by ~~Github~~.
- Added `--loglevel=` command-line option to the `mdbx_test` tool.
- Added few fast smoke-like tests into CMake builds.

Fixes:

- Fixed a race between starting a transaction and creating a DBI descriptor that could lead to `SIGSEGV` in the cursor tracking code.
- Clarified description of `MDBX_EPERM` error returned from `mdbx_env_set_geometry()`.
- Fixed non-promoting the parent transaction to be dirty in case the undo of the geometry update failed during abortion of a nested transaction.
- Resolved linking issues with `libstdc++fs`/`libc++fs`/`libc++experimental` for C++ `std::filesystem` or `std::experimental::filesystem` for legacy compilers.
- Added workaround for GNU Make 3.81 and earlier.
- Added workaround for Elbrus/LCC 1.25 compiler bug of class inline `static constexpr` member field.
- [Fixed](https://github.com/ledgerwatch/erigon/issues/3874) minor assertion regression (only debug builds were affected).
- Fixed detection of `C++20` concepts accessibility.
- Fixed detection of Clang's LTO availability for Android.
- Fixed extra definition of `_FILE_OFFSET_BITS=64` for Android that is problematic for 32-bit Bionic.
- Fixed build for ARM/ARM64 by MSVC.
- Fixed non-x86 Windows builds with `MDBX_WITHOUT_MSVC_CRT=ON` and `MDBX_BUILD_SHARED_LIBRARY=ON`.

Minors:

- Resolve minor MSVC warnings: avoid `/INCREMENTAL[:YES]` with `/LTCG`, `/W4` with `/W3`, the `C5105` warning.
- Switched to using `MDBX_EPERM` instead of `MDBX_RESULT_TRUE` to indicate that the geometry cannot be updated.
- Added `NULL` checking during memory allocation inside `mdbx_chk`.
- Resolved all warnings from MinGW while used without CMake.
- Added inheritable `target_include_directories()` to `CMakeLists.txt` for easy integration.
- Added build-time checks and paranoid runtime assertions for the `off_t` arguments of `fcntl()` which are used for locking.
- Added `-Wno-lto-type-mismatch` to avoid false-positive warnings from old GCC during LTO-enabled builds.
- Added checking for TID (system thread id) to avoid hang on 32-bit Bionic/Android within `pthread_mutex_lock()`.
- Reworked `MDBX_BUILD_TARGET` of CMake builds.
- Added `CMAKE_HOST_ARCH` and `CMAKE_HOST_CAN_RUN_EXECUTABLES_BUILT_FOR_TARGET`.


--------------------------------------------------------------------------------


## v0.11.6 at 2022-03-24

The stable release with the complete workaround for an incoherence flaw of Linux unified page/buffer cache.
Nonetheless the cause for this trouble may be an issue of Intel CPU cache/MESI.
See [issue#269](https://libmdbx.dqdkfa.ru/dead-github/issues/269) for more information.

Acknowledgments:

- [David Bouyssié](https://github.com/david-bouyssie) for [Scala bindings](https://github.com/david-bouyssie/mdbx4s).
- [Michelangelo Riccobene](https://github.com/mriccobene) for reporting and testing.

Fixes:

- [Added complete workaround](https://libmdbx.dqdkfa.ru/dead-github/issues/269) for an incoherence flaw of Linux unified page/buffer cache.
- [Fixed](https://libmdbx.dqdkfa.ru/dead-github/issues/272) cursor reusing for read-only transactions.
- Fixed copy&paste typo inside `mdbx::cursor::find_multivalue()`.

Minors:

- Minor refine C++ API for convenience.
- Minor internals refines.
- Added `lib-static` and `lib-shared` targets for make.
- Added minor workaround for AppleClang 13.3 bug.
- Clarified error messages of a signature/version mismatch.


--------------------------------------------------------------------------------


## v0.11.5 at 2022-02-23

The release with the temporary hotfix for a flaw of Linux unified page/buffer cache.
See [issue#269](https://libmdbx.dqdkfa.ru/dead-github/issues/269) for more information.

Acknowledgments:

- [Simon Leier](https://github.com/leisim) for reporting and testing.
- [Kai Wetlesen](https://github.com/kaiwetlesen) for [RPMs](http://copr.fedorainfracloud.org/coprs/kwetlesen/libmdbx/).
- [Tullio Canepa](https://github.com/canepat) for reporting C++ API issue and contributing.

Fixes:

- [Added hotfix](https://libmdbx.dqdkfa.ru/dead-github/issues/269) for a flaw of Linux unified page/buffer cache.
- [Fixed/Reworked](https://libmdbx.dqdkfa.ru/dead-github/pull/270) move-assignment operators for "managed" classes of C++ API.
- Fixed potential `SIGSEGV` while open DB with overrided non-default page size.
- [Made](https://libmdbx.dqdkfa.ru/dead-github/issues/267) `mdbx_env_open()` idempotence in failure cases.
- Refined/Fixed pages reservation inside `mdbx_update_gc()` to avoid non-reclamation in a rare cases.
- Fixed typo in a retained space calculation for the hsr-callback.

Minors:

- Reworked functions for meta-pages, split-off non-volatile.
- Disentangled C11-atomic fences/barriers and pure-functions (with `__attribute__((__pure__))`) to avoid compiler misoptimization.
- Fixed hypotetic unaligned access to 64-bit dwords on ARM with `__ARM_FEATURE_UNALIGNED` defined.
- Reasonable paranoia that makes clarity for code readers.
- Minor fixes Doxygen references, comments, descriptions, etc.


--------------------------------------------------------------------------------


## v0.11.4 at 2022-02-02

The stable release with fixes for large and huge databases sized of 4..128 TiB.

Acknowledgments:

- [Ledgerwatch](https://github.com/ledgerwatch), [Binance](https://github.com/binance-chain) and [Positive Technologies](https://www.ptsecurity.com/) teams for reporting, assistance in investigation and testing.
- [Alex Sharov](https://github.com/AskAlexSharov) for reporting, testing and provide resources for remote debugging/investigation.
- [Kris Zyp](https://github.com/kriszyp) for [Deno](https://deno.land/) support.

New features, extensions and improvements:

- Added treating the `UINT64_MAX` value as maximum for given option inside `mdbx_env_set_option()`.
- Added `to_hex/to_base58/to_base64::output(std::ostream&)` overloads without using temporary string objects as buffers.
- Added `--geometry-jitter=YES|no` option to the test framework.
- Added support for [Deno](https://deno.land/) support by [Kris Zyp](https://github.com/kriszyp).

Fixes:

- Fixed handling `MDBX_opt_rp_augment_limit` for GC's records from huge transactions (Erigon/Akula/Ethereum).
- [Fixed](https://libmdbx.dqdkfa.ru/dead-github/issues/258) build on Android (avoid including `sys/sem.h`).
- [Fixed](https://libmdbx.dqdkfa.ru/dead-github/pull/261) missing copy assignment operator for `mdbx::move_result`.
- Fixed missing `&` for `std::ostream &operator<<()` overloads.
- Fixed unexpected `EXDEV` (Cross-device link) error from `mdbx_env_copy()`.
- Fixed base64 encoding/decoding bugs in auxillary C++ API.
- Fixed overflow of `pgno_t` during checking PNL on 64-bit platforms.
- [Fixed](https://libmdbx.dqdkfa.ru/dead-github/issues/260) excessive PNL checking after sort for spilling.
- Reworked checking `MAX_PAGENO` and DB upper-size geometry limit.
- [Fixed](https://libmdbx.dqdkfa.ru/dead-github/issues/265) build for some combinations of versions of  MSVC and Windows SDK.

Minors:

- Added workaround for CLANG bug [D79919/PR42445](https://reviews.llvm.org/D79919).
- Fixed build test on Android (using `pthread_barrier_t` stub).
- Disabled C++20 concepts for CLANG < 14 on Android.
- Fixed minor `unused parameter` warning.
- Added CI for Android.
- Refine/cleanup internal logging.
- Refined line splitting inside hex/base58/base64 encoding to avoid `\n` at the end.
- Added workaround for modern libstdc++ with CLANG < 4.x
- Relaxed txn-check rules for auxiliary functions.
- Clarified a comments and descriptions, etc.
- Using the `-fno-semantic interposition` option to reduce the overhead to calling self own public functions.


--------------------------------------------------------------------------------


## v0.11.3 at 2021-12-31

Acknowledgments:

- [gcxfd <i@rmw.link>](https://github.com/gcxfd) for reporting, contributing and testing.
- [장세연 (Чан Се Ен)](https://github.com/sasgas) for reporting and testing.
- [Alex Sharov](https://github.com/AskAlexSharov) for reporting, testing and provide resources for remote debugging/investigation.

New features, extensions and improvements:

- [Added](https://libmdbx.dqdkfa.ru/dead-github/issues/236) `mdbx_cursor_get_batch()`.
- [Added](https://libmdbx.dqdkfa.ru/dead-github/issues/250) `MDBX_SET_UPPERBOUND`.
- C++ API is finalized now.
- The GC update stage has been [significantly speeded](https://libmdbx.dqdkfa.ru/dead-github/issues/254) when fixing huge Erigon's transactions (Ethereum ecosystem).

Fixes:

- Disabled C++20 concepts for stupid AppleClang 13.x
- Fixed internal collision of `MDBX_SHRINK_ALLOWED` with `MDBX_ACCEDE`.

Minors:

- Fixed returning `MDBX_RESULT_TRUE` (unexpected -1) from `mdbx_env_set_option()`.
- Added `mdbx_env_get_syncbytes()` and `mdbx_env_get_syncperiod()`.
- [Clarified](https://libmdbx.dqdkfa.ru/dead-github/pull/249) description of `MDBX_INTEGERKEY`.
- Reworked/simplified `mdbx_env_sync_internal()`.
- [Fixed](https://libmdbx.dqdkfa.ru/dead-github/issues/248) extra assertion inside `mdbx_cursor_put()` for `MDBX_DUPFIXED` cases.
- Avoiding extra looping inside `mdbx_env_info_ex()`.
- Explicitly enabled core dumps from stochastic tests scripts on Linux.
- [Fixed](https://libmdbx.dqdkfa.ru/dead-github/issues/253) `mdbx_override_meta()` to avoid false-positive assertions.
- For compatibility reverted returning `MDBX_ENODATA`for some cases.


--------------------------------------------------------------------------------


## v0.11.2 at 2021-12-02

Acknowledgments:

- [장세연 (Чан Се Ен)](https://github.com/sasgas) for contributing to C++ API.
- [Alain Picard](https://github.com/castortech) for [Java bindings](https://github.com/castortech/mdbxjni).
- [Alex Sharov](https://github.com/AskAlexSharov) for reporting and testing.
- [Kris Zyp](https://github.com/kriszyp) for reporting and testing.
- [Artem Vorotnikov](https://github.com/vorot93) for support [Rust wrapper](https://github.com/vorot93/libmdbx-rs).

Fixes:

- [Fixed compilation](https://libmdbx.dqdkfa.ru/dead-github/pull/239) with `devtoolset-9` on CentOS/RHEL 7.
- [Fixed unexpected `MDBX_PROBLEM` error](https://libmdbx.dqdkfa.ru/dead-github/issues/242) because of update an obsolete meta-page.
- [Fixed returning `MDBX_NOTFOUND` error](https://libmdbx.dqdkfa.ru/dead-github/issues/243) in case an inexact value found for `MDBX_GET_BOTH` operation.
- [Fixed compilation](https://libmdbx.dqdkfa.ru/dead-github/issues/245) without kernel/libc-devel headers.

Minors:

- Fixed `constexpr`-related macros for legacy compilers.
- Allowed to define 'CMAKE_CXX_STANDARD` using an environment variable.
- Simplified collection statistics of page operation .
- Added `MDBX_FORCE_BUILD_AS_MAIN_PROJECT` cmake option.
- Remove unneeded `#undef P_DIRTY`.


--------------------------------------------------------------------------------


## v0.11.1 at 2021-10-23

### Backward compatibility break:

The database format signature has been changed to prevent
forward-interoperability with an previous releases, which may lead to a
[false positive diagnosis of database corruption](https://libmdbx.dqdkfa.ru/dead-github/issues/238)
due to flaws of an old library versions.

This change is mostly invisible:

- previously versions are unable to read/write a new DBs;
- but the new release is able to handle an old DBs and will silently upgrade ones.

Acknowledgments:

- [Alex Sharov](https://github.com/AskAlexSharov) for reporting and testing.


********************************************************************************


## v0.10.5 at 2021-10-13 (obsolete, please use v0.11.1)

Unfortunately, the `v0.10.5` accidentally comes not full-compatible with previous releases:

- `v0.10.5` can read/processing DBs created by previous releases, i.e. the backward-compatibility is provided;
- however, previous releases may lead to false-corrupted state with DB that was touched by `v0.10.5`, i.e. the forward-compatibility is broken for `v0.10.4` and earlier.

This cannot be fixed, as it requires fixing past versions, which as a result we will just get a current version.
Therefore, it is recommended to use `v0.11.1` instead of `v0.10.5`.

Acknowledgments:

- [Noel Kuntze](https://github.com/Thermi) for immediately bug reporting.

Fixes:

- Fixed unaligned access regression after the `#pragma pack` fix for modern compilers.
- Added UBSAN-test to CI to avoid a regression(s) similar to lately fixed.
- Fixed possibility of meta-pages clashing after manually turn to a particular meta-page using `mdbx_chk` utility.

Minors:

- Refined handling of weak or invalid meta-pages while a DB opening.
- Refined providing information for the `@MAIN` and `@GC` sub-databases of a last committed modification transaction's ID.


--------------------------------------------------------------------------------


## v0.10.4 at 2021-10-10

Acknowledgments:

- [Artem Vorotnikov](https://github.com/vorot93) for support [Rust wrapper](https://github.com/vorot93/libmdbx-rs).
- [Andrew Ashikhmin](https://github.com/yperbasis) for contributing to C++ API.

Fixes:

- Fixed possibility of looping update GC during transaction commit (no public issue since the problem was discovered inside [Positive Technologies](https://www.ptsecurity.ru)).
- Fixed `#pragma pack` to avoid provoking some compilers to generate code with [unaligned access](https://libmdbx.dqdkfa.ru/dead-github/issues/235).
- Fixed `noexcept` for potentially throwing `txn::put()` of C++ API.

Minors:

- Added stochastic test script for checking small transactions cases.
- Removed extra transaction commit/restart inside test framework.
- In debugging builds fixed a too small (single page) by default DB shrink threshold.


--------------------------------------------------------------------------------


## v0.10.3 at 2021-08-27

Acknowledgments:

- [Francisco Vallarino](https://github.com/fjvallarino) for [Haskell bindings for libmdbx](https://hackage.haskell.org/package/libmdbx).
- [Alex Sharov](https://github.com/AskAlexSharov) for reporting and testing.
- [Andrea Lanfranchi](https://github.com/AndreaLanfranchi) for contributing.

Extensions and improvements:

- Added `cursor::erase()` overloads for `key` and for `key-value`.
- Resolve minor Coverity Scan issues (no fixes but some hint/comment were added).
- Resolve minor UndefinedBehaviorSanitizer issues (no fixes but some workaround were added).

Fixes:

- Always setup `madvise` while opening DB (fixes https://libmdbx.dqdkfa.ru/dead-github/issues/231).
- Fixed checking legacy `P_DIRTY` flag (`0x10`) for nested/sub-pages.

Minors:

- Fixed getting revision number from middle of history during amalgamation (GNU Makefile).
- Fixed search GCC tools for LTO (CMake scripts).
- Fixed/reorder dirs list for search CLANG tools for LTO (CMake scripts).
- Fixed/workarounds for CLANG < 9.x
- Fixed CMake warning about compatibility with 3.8.2


--------------------------------------------------------------------------------


## v0.10.2 at 2021-07-26

Acknowledgments:

- [Alex Sharov](https://github.com/AskAlexSharov) for reporting and testing.
- [Andrea Lanfranchi](https://github.com/AndreaLanfranchi) for reporting bugs.
- [Lionel Debroux](https://github.com/debrouxl) for fuzzing tests and reporting bugs.
- [Sergey Fedotov](https://github.com/SergeyFromHell/) for [`node-mdbx` NodeJS bindings](https://www.npmjs.com/package/node-mdbx).
- [Kris Zyp](https://github.com/kriszyp) for [`lmdbx-store` NodeJS bindings](https://github.com/kriszyp/lmdbx-store).
- [Noel Kuntze](https://github.com/Thermi) for [draft Python bindings](https://libmdbx.dqdkfa.ru/dead-github/commits/python-bindings).

New features, extensions and improvements:

- [Allow to predefine/override `MDBX_BUILD_TIMESTAMP` for builds reproducibility](https://libmdbx.dqdkfa.ru/dead-github/issues/201).
- Added options support for `long-stochastic` script.
- Avoided `MDBX_TXN_FULL` error for large transactions when possible.
- The `MDBX_READERS_LIMIT` increased to `32767`.
- Raise `MDBX_TOO_LARGE` under Valgrind/ASAN if being opened DB is 100 larger than RAM (to avoid hangs and OOM).
- Minimized the size of poisoned/unpoisoned regions to avoid Valgrind/ASAN stuck.
- Added more workarounds for QEMU for testing builds for 32-bit platforms, Alpha and Sparc architectures.
- `mdbx_chk` now skips iteration & checking of DB' records if corresponding page-tree is corrupted (to avoid `SIGSEGV`, ASAN failures, etc).
- Added more checks for [rare/fuzzing corruption cases](https://libmdbx.dqdkfa.ru/dead-github/issues/217).

Backward compatibility break:

- Use file `VERSION.txt` for version information instead of `VERSION` to avoid collision with `#include <version>`.
- Rename `slice::from/to_FOO_bytes()` to `slice::envisage_from/to_FOO_length()'.
- Rename `MDBX_TEST_EXTRA` make's variable to `MDBX_SMOKE_EXTRA`.
- Some details of the C++ API have been changed for subsequent freezing.

Fixes:

- Fixed excess meta-pages checks in case `mdbx_chk` is called to check the DB for a specific meta page and thus could prevent switching to the selected meta page, even if the check passed without errors.
- Fixed [recursive use of SRW-lock on Windows cause by `MDBX_NOTLS` option](https://libmdbx.dqdkfa.ru/dead-github/issues/203).
- Fixed [log a warning during a new DB creation](https://libmdbx.dqdkfa.ru/dead-github/issues/205).
- Fixed [false-negative `mdbx_cursor_eof()` result](https://libmdbx.dqdkfa.ru/dead-github/issues/207).
- Fixed [`make install` with non-GNU `install` utility (OSX, BSD)](https://libmdbx.dqdkfa.ru/dead-github/issues/208).
- Fixed [installation by `CMake` in special cases by complete use `GNUInstallDirs`'s variables](https://libmdbx.dqdkfa.ru/dead-github/issues/209).
- Fixed [C++ Buffer issue with `std::string` and alignment](https://libmdbx.dqdkfa.ru/dead-github/issues/191).
- Fixed `safe64_reset()` for platforms without atomic 64-bit compare-and-swap.
- Fixed hang/shutdown on big-endian platforms without `__cxa_thread_atexit()`.
- Fixed [using bad meta-pages if DB was partially/recoverable corrupted](https://libmdbx.dqdkfa.ru/dead-github/issues/217).
- Fixed extra `noexcept` for `buffer::&assign_reference()`.
- Fixed `bootid` generation on Windows for case of change system' time.
- Fixed [test framework keygen-related issue](https://libmdbx.dqdkfa.ru/dead-github/issues/127).


--------------------------------------------------------------------------------


## v0.10.1 at 2021-06-01

Acknowledgments:

- [Alexey Akhunov](https://github.com/AlexeyAkhunov) and [Alex Sharov](https://github.com/AskAlexSharov) for bug reporting and testing.
- [Andrea Lanfranchi](https://github.com/AndreaLanfranchi) for bug reporting and testing related to WSL2.

New features:

- Added `-p` option to `mdbx_stat` utility for printing page operations statistic.
- Added explicit checking for and warning about using unfit github's archives.
- Added fallback from [OFD locking](https://bit.ly/3yFRtYC) to legacy non-OFD POSIX file locks on an `EINVAL` error.
- Added [Plan 9](https://en.wikipedia.org/wiki/9P_(protocol)) network file system to the whitelist for an ability to open a DB in exclusive mode.
- Support for opening from WSL2 environment a DB hosted on Windows drive and mounted via [DrvFs](https://docs.microsoft.com/it-it/archive/blogs/wsl/wsl-file-system-support#drvfs) (i.e by Plan 9 noted above).

Fixes:

- Fixed minor "foo not used" warnings from modern C++ compilers when building the C++ part of the library.
- Fixed confusing/messy errors when build library from unfit github's archives (https://libmdbx.dqdkfa.ru/dead-github/issues/197).
- Fixed `#​e​l​s​i​f` typo.
- Fixed rare unexpected `MDBX_PROBLEM` error during altering data in huge transactions due to wrong spilling/oust of dirty pages (https://libmdbx.dqdkfa.ru/dead-github/issues/195).
- Re-Fixed WSL1/WSL2 detection with distinguishing (https://libmdbx.dqdkfa.ru/dead-github/issues/97).


--------------------------------------------------------------------------------


## v0.10.0 at 2021-05-09

Acknowledgments:

- [Mahlon E. Smith](https://github.com/mahlonsmith) for [Ruby bindings](https://rubygems.org/gems/mdbx/).
- [Alex Sharov](https://github.com/AskAlexSharov) for [mdbx-go](https://github.com/torquem-ch/mdbx-go), bug reporting and testing.
- [Artem Vorotnikov](https://github.com/vorot93) for bug reporting and PR.
- [Paolo Rebuffo](https://www.linkedin.com/in/paolo-rebuffo-8255766/), [Alexey Akhunov](https://github.com/AlexeyAkhunov) and Mark Grosberg for donations.
- [Noel Kuntze](https://github.com/Thermi) for preliminary [Python bindings](https://github.com/Thermi/libmdbx/tree/python-bindings)

New features:

- Added `mdbx_env_set_option()` and `mdbx_env_get_option()` for controls
  various runtime options for an environment (announce of this feature  was missed in a previous news).
- Added `MDBX_DISABLE_PAGECHECKS` build option to disable some checks to reduce an overhead
  and detection probability of database corruption to a values closer to the LMDB.
  The `MDBX_DISABLE_PAGECHECKS=1` provides a performance boost of about 10% in CRUD scenarios,
  and conjointly with the `MDBX_ENV_CHECKPID=0` and `MDBX_TXN_CHECKOWNER=0` options can yield
  up to 30% more performance compared to LMDB.
- Using float point (exponential quantized) representation for internal 16-bit values
  of grow step and shrink threshold when huge ones (https://libmdbx.dqdkfa.ru/dead-github/issues/166).
  To minimize the impact on compatibility, only the odd values inside the upper half
  of the range (i.e. 32769..65533) are used for the new representation.
- Added the `mdbx_drop` similar to LMDB command-line tool to purge or delete (sub)database(s).
- [Ruby bindings](https://rubygems.org/gems/mdbx/) is available now by [Mahlon E. Smith](https://github.com/mahlonsmith).
- Added `MDBX_ENABLE_MADVISE` build option which controls the use of POSIX `madvise()` hints and friends.
- The internal node sizes were refined, resulting in a reduction in large/overflow pages in some use cases
  and a slight increase in limits for a keys size to ≈½ of page size.
- Added to `mdbx_chk` output number of keys/items on pages.
- Added explicit `install-strip` and `install-no-strip` targets to the `Makefile` (https://libmdbx.dqdkfa.ru/dead-github/pull/180).
- Major rework page splitting (af9b7b560505684249b76730997f9e00614b8113) for
   - An "auto-appending" feature upon insertion for both ascending and
     descending key sequences. As a result, the optimality of page filling
     increases significantly (more densely, less slackness) while
     inserting ordered sequences of keys,
   - A "splitting at middle" to make page tree more balanced on average.
- Added `mdbx_get_sysraminfo()` to the API.
- Added guessing a reasonable maximum DB size for the default upper limit of geometry (https://libmdbx.dqdkfa.ru/dead-github/issues/183).
- Major rework internal labeling of a dirty pages (958fd5b9479f52f2124ab7e83c6b18b04b0e7dda) for
  a "transparent spilling" feature with the gist to make a dirty pages
  be ready to spilling (writing to a disk) without further altering ones.
  Thus in the `MDBX_WRITEMAP` mode the OS kernel able to oust dirty pages
  to DB file without further penalty during transaction commit.
  As a result, page swapping and I/O could be significantly reduced during extra large transactions and/or lack of memory.
- Minimized reading leaf-pages during dropping subDB(s) and nested trees.
- Major rework a spilling of dirty pages to support [LRU](https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU))
  policy and prioritization for a large/overflow pages.
- Statistics of page operations (split, merge, copy, spill, etc) now available through `mdbx_env_info_ex()`.
- Auto-setup limit for length of dirty pages list (`MDBX_opt_txn_dp_limit` option).
- Support `make options` to list available build options.
- Support `make help` to list available make targets.
- Silently `make`'s build by default.
- Preliminary [Python bindings](https://github.com/Thermi/libmdbx/tree/python-bindings) is available now
  by [Noel Kuntze](https://github.com/Thermi) (https://libmdbx.dqdkfa.ru/dead-github/issues/147).

Backward compatibility break:

- The `MDBX_AVOID_CRT` build option was renamed to `MDBX_WITHOUT_MSVC_CRT`.
  This option is only relevant when building for Windows.
- The `mdbx_env_stat()` always, and `mdbx_env_stat_ex()` when called with the zeroed transaction parameter,
  now internally start temporary read transaction and thus may returns `MDBX_BAD_RSLOT` error.
  So, just never use deprecated `mdbx_env_stat()' and call `mdbx_env_stat_ex()` with transaction parameter.
- The build option `MDBX_CONFIG_MANUAL_TLS_CALLBACK` was removed and now just a non-zero value of
  the `MDBX_MANUAL_MODULE_HANDLER` macro indicates the requirement to manually call `mdbx_module_handler()`
  when loading libraries and applications uses statically linked libmdbx on an obsolete Windows versions.

Fixes:

- Fixed performance regression due non-optimal C11 atomics usage (https://libmdbx.dqdkfa.ru/dead-github/issues/160).
- Fixed "reincarnation" of subDB after it deletion (https://libmdbx.dqdkfa.ru/dead-github/issues/168).
- Fixed (disallowing) implicit subDB deletion via operations on `@MAIN`'s DBI-handle.
- Fixed a crash of `mdbx_env_info_ex()` in case of a call for a non-open environment (https://libmdbx.dqdkfa.ru/dead-github/issues/171).
- Fixed the selecting/adjustment values inside `mdbx_env_set_geometry()` for implicit out-of-range cases (https://libmdbx.dqdkfa.ru/dead-github/issues/170).
- Fixed `mdbx_env_set_option()` for set initial and limit size of dirty page list ((https://libmdbx.dqdkfa.ru/dead-github/issues/179).
- Fixed an unreasonably huge default upper limit for DB geometry (https://libmdbx.dqdkfa.ru/dead-github/issues/183).
- Fixed `constexpr` specifier for the `slice::invalid()`.
- Fixed (no)readahead auto-handling (https://libmdbx.dqdkfa.ru/dead-github/issues/164).
- Fixed non-alloy build for Windows.
- Switched to using Heap-functions instead of LocalAlloc/LocalFree on Windows.
- Fixed `mdbx_env_stat_ex()` to returning statistics of the whole environment instead of MainDB only (https://libmdbx.dqdkfa.ru/dead-github/issues/190).
- Fixed building by GCC 4.8.5 (added workaround for a preprocessor's bug).
- Fixed building C++ part for iOS <= 13.0 (unavailability of  `std::filesystem::path`).
- Fixed building for Windows target versions prior to Windows Vista (`WIN32_WINNT < 0x0600`).
- Fixed building by MinGW for Windows (https://libmdbx.dqdkfa.ru/dead-github/issues/155).


********************************************************************************


## v0.9.3 at 2021-02-02

Acknowledgments:

- [Mahlon E. Smith](http://www.martini.nu/) for [FreeBSD port of libmdbx](https://svnweb.freebsd.org/ports/head/databases/mdbx/).
- [장세연](http://www.castis.com) for bug fixing and PR.
- [Clément Renault](https://github.com/Kerollmops/heed) for [Heed](https://github.com/Kerollmops/heed) fully typed Rust wrapper.
- [Alex Sharov](https://github.com/AskAlexSharov) for bug reporting.
- [Noel Kuntze](https://github.com/Thermi) for bug reporting.

Removed options and features:

- Drop `MDBX_HUGE_TRANSACTIONS` build-option (now no longer required).

New features:

- Package for FreeBSD is available now by Mahlon E. Smith.
- New API functions to get/set various options (https://libmdbx.dqdkfa.ru/dead-github/issues/128):
   - the maximum number of named databases for the environment;
   - the maximum number of threads/reader slots;
   - threshold (since the last unsteady commit) to force flush the data buffers to disk;
   - relative period (since the last unsteady commit) to force flush the data buffers to disk;
   - limit to grow a list of reclaimed/recycled page's numbers for finding a sequence of contiguous pages for large data items;
   - limit to grow a cache of dirty pages for reuse in the current transaction;
   - limit of a pre-allocated memory items for dirty pages;
   - limit of dirty pages for a write transaction;
   - initial allocation size for dirty pages list of a write transaction;
   - maximal part of the dirty pages may be spilled when necessary;
   - minimal part of the dirty pages should be spilled when necessary;
   - how much of the parent transaction dirty pages will be spilled while start each child transaction;
- Unlimited/Dynamic size of retired and dirty page lists (https://libmdbx.dqdkfa.ru/dead-github/issues/123).
- Added `-p` option (purge subDB before loading) to `mdbx_load` tool.
- Reworked spilling of large transaction and committing of nested transactions:
   - page spilling code reworked to avoid the flaws and bugs inherited from LMDB;
   - limit for number of dirty pages now is controllable at runtime;
   - a spilled pages, including overflow/large pages, now can be reused and refunded/compactified in nested transactions;
   - more effective refunding/compactification especially for the loosed page cache.
- Added `MDBX_ENABLE_REFUND` and `MDBX_PNL_ASCENDING` internal/advanced build options.
- Added `mdbx_default_pagesize()` function.
- Better support architectures with a weak/relaxed memory consistency model (ARM, AARCH64, PPC, MIPS, RISC-V, etc) by means [C11 atomics](https://en.cppreference.com/w/c/atomic).
- Speed up page number lists and dirty page lists (https://libmdbx.dqdkfa.ru/dead-github/issues/132).
- Added `LIBMDBX_NO_EXPORTS_LEGACY_API` build option.

Fixes:

- Fixed missing cleanup (null assigned) in the C++ commit/abort (https://libmdbx.dqdkfa.ru/dead-github/pull/143).
- Fixed `mdbx_realloc()` for case of nullptr and `MDBX_WITHOUT_MSVC_CRT=ON` for Windows.
- Fixed the possibility to use invalid and renewed (closed & re-opened, dropped & re-created) DBI-handles (https://libmdbx.dqdkfa.ru/dead-github/issues/146).
- Fixed 4-byte aligned access to 64-bit integers, including access to the `bootid` meta-page's field (https://libmdbx.dqdkfa.ru/dead-github/issues/153).
- Fixed minor/potential memory leak during page flushing and unspilling.
- Fixed handling states of cursors's and subDBs's for nested transactions.
- Fixed page leak in extra rare case the list of retired pages changed during update GC on transaction commit.
- Fixed assertions to avoid false-positive UB detection by CLANG/LLVM (https://libmdbx.dqdkfa.ru/dead-github/issues/153).
- Fixed `MDBX_TXN_FULL` and regressive `MDBX_KEYEXIST` during large transaction commit with `MDBX_LIFORECLAIM` (https://libmdbx.dqdkfa.ru/dead-github/issues/123).
- Fixed auto-recovery (`weak->steady` with the same boot-id) when Database size at last weak checkpoint is large than at last steady checkpoint.
- Fixed operation on systems with unusual small/large page size, including PowerPC (https://libmdbx.dqdkfa.ru/dead-github/issues/157).


--------------------------------------------------------------------------------


## v0.9.2 at 2020-11-27

Acknowledgments:

- Jens Alfke (Mobile Architect at [Couchbase](https://www.couchbase.com/)) for [NimDBX](https://github.com/snej/nimdbx).
- Clément Renault (CTO at [MeiliSearch](https://www.meilisearch.com/)) for [mdbx-rs](https://github.com/Kerollmops/mdbx-rs).
- Alex Sharov (Go-Lang Teach Lead at [TurboGeth/Ethereum](https://ethereum.org/)) for an extreme test cases and bug reporting.
- George Hazan (CTO at [Miranda NG](https://www.miranda-ng.org/)) for bug reporting.
- [Positive Technologies](https://www.ptsecurity.com/) for funding and [The Standoff](https://standoff365.com/).

Added features:

- Provided package for [buildroot](https://buildroot.org/).
- Binding for Nim is [available](https://github.com/snej/nimdbx) now by Jens Alfke.
- Added `mdbx_env_delete()` for deletion an environment files in a proper and multiprocess-safe way.
- Added `mdbx_txn_commit_ex()` with collecting latency information.
- Fast completion pure nested transactions.
- Added `LIBMDBX_INLINE_API` macro and inline versions of some API functions.
- Added `mdbx_cursor_copy()` function.
- Extended tests for checking cursor tracking.
- Added `MDBX_SET_LOWERBOUND` operation for `mdbx_cursor_get()`.

Fixes:

- Fixed missing installation of `mdbx.h++`.
- Fixed use of obsolete `__noreturn`.
- Fixed use of `yield` instruction on ARM if unsupported.
- Added pthread workaround for buggy toolchain/cmake/buildroot.
- Fixed use of `pthread_yield()` for non-GLIBC.
- Fixed use of `RegGetValueA()` on Windows 2000/XP.
- Fixed use of `GetTickCount64()` on Windows 2000/XP.
- Fixed opening DB on a network shares (in the exclusive mode).
- Fixed copy&paste typos.
- Fixed minor false-positive GCC warning.
- Added workaround for broken `DEFINE_ENUM_FLAG_OPERATORS` from Windows SDK.
- Fixed cursor state after multimap/dupsort repeated deletes (https://libmdbx.dqdkfa.ru/dead-github/issues/121).
- Added `SIGPIPE` suppression for internal thread during `mdbx_env_copy()`.
- Fixed extra-rare `MDBX_KEY_EXIST` error during `mdbx_commit()` (https://libmdbx.dqdkfa.ru/dead-github/issues/131).
- Fixed spilled pages checking (https://libmdbx.dqdkfa.ru/dead-github/issues/126).
- Fixed `mdbx_load` for 'plain text' and without `-s name` cases (https://libmdbx.dqdkfa.ru/dead-github/issues/136).
- Fixed save/restore/commit of cursors for nested transactions.
- Fixed cursors state in rare/special cases (move next beyond end-of-data, after deletion and so on).
- Added workaround for MSVC 19.28 (Visual Studio 16.8) (but may still hang during compilation).
- Fixed paranoidal Clang C++ UB for bitwise operations with flags defined by enums.
- Fixed large pages checking (for compatibility and to avoid false-positive errors from `mdbx_chk`).
- Added workaround for Wine (https://github.com/miranda-ng/miranda-ng/issues/1209).
- Fixed `ERROR_NOT_SUPPORTED` while opening DB by UNC pathnames (https://github.com/miranda-ng/miranda-ng/issues/2627).
- Added handling `EXCEPTION_POSSIBLE_DEADLOCK` condition for Windows.


--------------------------------------------------------------------------------


## v0.9.1 2020-09-30

Added features:

- Preliminary C++ API with support for C++17 polymorphic allocators.
- [Online C++ API reference](https://libmdbx.dqdkfa.ru/) by Doxygen.
- Quick reference for Insert/Update/Delete operations.
- Explicit `MDBX_SYNC_DURABLE` to sync modes for API clarity.
- Explicit `MDBX_ALLDUPS` and `MDBX_UPSERT` for API clarity.
- Support for read transactions preparation (`MDBX_TXN_RDONLY_PREPARE` flag).
- Support for cursor preparation/(pre)allocation and reusing (`mdbx_cursor_create()` and `mdbx_cursor_bind()` functions).
- Support for checking database using specified meta-page (see `mdbx_chk -h`).
- Support for turn to the specific meta-page after checking (see `mdbx_chk -h`).
- Support for explicit reader threads (de)registration.
- The `mdbx_txn_break()` function to explicitly mark a transaction as broken.
- Improved handling of corrupted databases by `mdbx_chk` utility and `mdbx_walk_tree()` function.
- Improved DB corruption detection by checking parent-page-txnid.
- Improved opening large DB (> 4Gb) from 32-bit code.
- Provided `pure-function` and `const-function` attributes to C API.
- Support for user-settable context for transactions & cursors.
- Revised API and documentation related to Handle-Slow-Readers callback feature.

Deprecated functions and flags:

- For clarity and API simplification the `MDBX_MAPASYNC` flag is deprecated.
  Just use `MDBX_SAFE_NOSYNC` or `MDBX_UTTERLY_NOSYNC` instead of it.
- `MDBX_oom_func`, `mdbx_env_set_oomfunc()` and `mdbx_env_get_oomfunc()`
  replaced with `MDBX_hsr_func`, `mdbx_env_get_hsr` and `mdbx_env_get_hsr()`.

Fixes:

- Fix `mdbx_strerror()` for `MDBX_BUSY` error (no error description is returned).
- Fix update internal meta-geo information in read-only mode (`EACCESS` or `EBADFD` error).
- Fix `mdbx_page_get()` null-defer when DB corrupted (crash by `SIGSEGV`).
- Fix `mdbx_env_open()` for re-opening after non-fatal errors (`mdbx_chk` unexpected failures).
- Workaround for MSVC 19.27 `static_assert()` bug.
- Doxygen descriptions and refinement.
- Update Valgrind's suppressions.
- Workaround to avoid infinite loop of 'nested' testcase on MIPS under QEMU.
- Fix a lot of typos & spelling (Thanks to Josh Soref for PR).
- Fix `getopt()` messages for Windows (Thanks to Andrey Sporaw for reporting).
- Fix MSVC compiler version requirements (Thanks to Andrey Sporaw for reporting).
- Workarounds for QEMU's bugs to run tests for cross-built[A library under QEMU.
- Now C++ compiler optional for building by CMake.


--------------------------------------------------------------------------------


## v0.9.0 2020-07-31 (not a release, but API changes)

Added features:

- [Online C API reference](https://libmdbx.dqdkfa.ru/) by Doxygen.
- Separated enums for environment, sub-databases, transactions, copying and data-update flags.

Deprecated functions and flags:

- Usage of custom comparators and the `mdbx_dbi_open_ex()` are deprecated, since such databases couldn't be checked by the `mdbx_chk` utility.
  Please use the value-to-key functions to provide keys that are compatible with the built-in libmdbx comparators.


********************************************************************************


## 2020-07-06

- Added support multi-opening the same DB in a process with SysV locking (BSD).
- Fixed warnings & minors for LCC compiler (E2K).
- Enabled to simultaneously open the same database from processes with and without the `MDBX_WRITEMAP` option.
- Added key-to-value, `mdbx_get_keycmp()` and `mdbx_get_datacmp()` functions (helpful to avoid using custom comparators).
- Added `ENABLE_UBSAN` CMake option to enabling the UndefinedBehaviorSanitizer from GCC/CLANG.
- Workaround for [CLANG bug](https://bugs.llvm.org/show_bug.cgi?id=43275).
- Returning `MDBX_CORRUPTED` in case all meta-pages are weak and no other error.
- Refined mode bits while auto-creating LCK-file.
- Avoids unnecessary database file re-mapping in case geometry changed by another process(es).
  From the user's point of view, the `MDBX_UNABLE_EXTEND_MAPSIZE` error will now be returned less frequently and only when using the DB in the current process really requires it to be reopened.
- Remapping on-the-fly and of the database file was implemented.
  Now remapping with a change of address is performed automatically if there are no dependent readers in the current process.


## 2020-06-12

- Minor change versioning. The last number in the version now means the number of commits since last release/tag.
- Provide ChangeLog file.
- Fix for using libmdbx as a C-only sub-project with CMake.
- Fix `mdbx_env_set_geometry()` for case it is called from an opened environment outside of a write transaction.
- Add support for huge transactions and `MDBX_HUGE_TRANSACTIONS` build-option (default `OFF`).
- Refine LTO (link time optimization) for clang.
- Force enabling exceptions handling for MSVC (`/EHsc` option).


## 2020-06-05

- Support for Android/Bionic.
- Support for iOS.
- Auto-handling `MDBX_NOSUBDIR` while opening for any existing database.
- Engage github-actions to make release-assets.
- Clarify API description.
- Extended keygen-cases in stochastic test.
- Fix fetching of first/lower key from LEAF2-page during page merge.
- Fix missing comma in array of error messages.
- Fix div-by-zero while copy-with-compaction for non-resizable environments.
- Fixes & enhancements for custom-comparators.
- Fix `MDBX_WITHOUT_MSVC_CRT` option and missing `ntdll.def`.
- Fix `mdbx_env_close()` to work correctly called concurrently from several threads.
- Fix null-deref in an ASAN-enabled builds while opening the environment with error and/or read-only.
- Fix AddressSanitizer errors after closing the environment.
- Fix/workaround to avoid GCC 10.x pedantic warnings.
- Fix using `ENODATA` for FreeBSD.
- Avoid invalidation of DBI-handle(s) when it just closes.
- Avoid using `pwritev()` for single-writes (up to 10% speedup for some kernels & scenarios).
- Avoiding `MDBX_UTTERLY_NOSYNC` as result of flags merge.
- Add `mdbx_dbi_dupsort_depthmask()` function.
- Add `MDBX_CP_FORCE_RESIZABLE` option.
- Add deprecated `MDBX_MAP_RESIZED` for compatibility.
- Add `MDBX_BUILD_TOOLS` option (default `ON`).
- Refine `mdbx_dbi_open_ex()` to safe concurrently opening the same handle from different threads.
- Truncate clk-file during environment closing. So a zero-length lck-file indicates that the environment was closed properly.
- Refine `mdbx_update_gc()` for huge transactions with small sizes of database page.
- Extends dump/load to support all MDBX attributes.
- Avoid upsertion the same key-value data, fix related assertions.
- Rework min/max length checking for keys & values.
- Checking the order of keys on all pages during checking.
- Support `CFLAGS_EXTRA` make-option for convenience.
- Preserve the last txnid while copying with compactification.
- Auto-reset running transaction in mdbx_txn_renew().
- Automatically abort errored transaction in mdbx_txn_commit().
- Auto-choose page size for large databases.
- Rearrange source files, rework build, options-support by CMake.
- Crutch for WSL1 (Windows subsystem for Linux).
- Refine install/uninstall targets.
- Support for Valgrind 3.14 and later.
- Add check-analyzer check-ubsan check-asan check-leak targets to Makefile.
- Minor fix/workaround to avoid UBSAN traps for `memcpy(ptr, NULL, 0)`.
- Avoid some GCC-analyzer false-positive warnings.


## 2020-03-18

- Workarounds for Wine (Windows compatibility layer for Linux).
- `MDBX_MAP_RESIZED` renamed to `MDBX_UNABLE_EXTEND_MAPSIZE`.
- Clarify API description, fix typos.
- Speedup runtime checks in debug/checked builds.
- Added checking for read/write transactions overlapping for the same thread, added `MDBX_TXN_OVERLAPPING` error and `MDBX_DBG_LEGACY_OVERLAP` option.
- Added `mdbx_key_from_jsonInteger()`, `mdbx_key_from_double()`, `mdbx_key_from_float()`, `mdbx_key_from_int64()` and `mdbx_key_from_int32()` functions. See `mdbx.h` for description.
- Fix compatibility (use zero for invalid DBI).
- Refine/clarify error messages.
- Avoids extra error messages "bad txn" from mdbx_chk when DB is corrupted.


## 2020-01-21

- Fix `mdbx_load` utility for custom comparators.
- Fix checks related to `MDBX_APPEND` flag inside `mdbx_cursor_put()`.
- Refine/fix dbi_bind() internals.
- Refine/fix handling `STATUS_CONFLICTING_ADDRESSES`.
- Rework `MDBX_DBG_DUMP` option to avoid disk I/O performance degradation.
- Add built-in help to test tool.
- Fix `mdbx_env_set_geometry()` for large page size.
- Fix env_set_geometry() for large pagesize.
- Clarify API description & comments, fix typos.


## 2019-12-31

- Fix returning MDBX_RESULT_TRUE from page_alloc().
- Fix false-positive ASAN issue.
- Fix assertion for `MDBX_NOTLS` option.
- Rework `MADV_DONTNEED` threshold.
- Fix `mdbx_chk` utility for don't checking some numbers if walking on the B-tree was disabled.
- Use page's mp_txnid for basic integrity checking.
- Add `MDBX_FORCE_ASSERTIONS` built-time option.
- Rework `MDBX_DBG_DUMP` to avoid performance degradation.
- Rename `MDBX_NOSYNC` to `MDBX_SAFE_NOSYNC` for clarity.
- Interpret `ERROR_ACCESS_DENIED` from `OpenProcess()` as 'process exists'.
- Avoid using `FILE_FLAG_NO_BUFFERING` for compatibility with small database pages.
- Added install section for CMake.


## 2019-12-02

- Support for Mac OSX, FreeBSD, NetBSD, OpenBSD, DragonFly BSD, OpenSolaris, OpenIndiana (AIX and HP-UX pending).
- Use bootid for decisions of rollback.
- Counting retired pages and extended transaction info.
- Add `MDBX_ACCEDE` flag for database opening.
- Using OFD-locks and tracking for in-process multi-opening.
- Hot backup into pipe.
- Support for cmake & amalgamated sources.
- Fastest internal sort implementation.
- New internal dirty-list implementation with lazy sorting.
- Support for lazy-sync-to-disk with polling.
- Extended key length.
- Last update transaction number for each sub-database.
- Automatic read ahead enabling/disabling.
- More auto-compactification.
- Using -fsanitize=undefined and -Wpedantic options.
- Rework page merging.
- Nested transactions.
- API description.
- Checking for non-local filesystems to avoid DB corruption.


********************************************************************************


For early changes see the git commit history.
