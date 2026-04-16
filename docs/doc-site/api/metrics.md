---
title: "Quick API index"
description: |
  API Quick Index & Metrics.
weight: -1
---

## Domains

All assertions are classified into **19** domains to help navigate the API, depending on your use case.

## API metrics

Counts for core functionality, excluding variants (formatted, forward, forward-formatted).

| Kind                     | Count             |
| ------------------------ | ----------------- |
| All functions            | 135  |
| All core assertions      | 131 |
| Generic assertions       | 49   |
| Helpers (not assertions) | 4    |
| Others                   | 0     |

## Quick index

Table of core assertions, excluding variants. Each function is side by side with its logical opposite (when available).

| Assertion                | Opposite          | Domain | Kind |
| ------------------------ | ----------------- | ------ | ---- |
| [CallerInfo](common/#callerinfo) |  | common | helper |
| [Condition](condition/#condition) |  | condition |  |
| [Consistently[C Conditioner]](condition/#consistentlyc-conditioner) {{% icon icon="star" color=orange %}} |  | condition |  |
| [Contains](collection/#contains) | [NotContains](collection/#notcontains) | collection |  |
| [DirExists](file/#direxists) | [DirNotExists](file/#dirnotexists) | file |  |
| [ElementsMatch](collection/#elementsmatch) | [NotElementsMatch](collection/#notelementsmatch) | collection |  |
| [ElementsMatchT[E comparable]](collection/#elementsmatchte-comparable) {{% icon icon="star" color=orange %}} | [NotElementsMatchT](collection/#notelementsmatchte-comparable) | collection |  |
| [Empty](equality/#empty) | [NotEmpty](equality/#notempty) | equality |  |
| [Equal](equality/#equal) | [NotEqual](equality/#notequal) | equality |  |
| [EqualError](error/#equalerror) |  | error |  |
| [EqualExportedValues](equality/#equalexportedvalues) |  | equality |  |
| [EqualT[V comparable]](equality/#equaltv-comparable) {{% icon icon="star" color=orange %}} | [NotEqualT](equality/#notequaltv-comparable) | equality |  |
| [EqualValues](equality/#equalvalues) | [NotEqualValues](equality/#notequalvalues) | equality |  |
| [Error](error/#error) | [NoError](error/#noerror) | error |  |
| [ErrorAs](error/#erroras) | [NotErrorAs](error/#noterroras) | error |  |
| [ErrorContains](error/#errorcontains) |  | error |  |
| [ErrorIs](error/#erroris) | [NotErrorIs](error/#noterroris) | error |  |
| [EventuallyWith[C CollectibleConditioner]](condition/#eventuallywithc-collectibleconditioner) {{% icon icon="star" color=orange %}} |  | condition |  |
| [Eventually[C Conditioner]](condition/#eventuallyc-conditioner) {{% icon icon="star" color=orange %}} | [Never](condition/#never) | condition |  |
| [Exactly](equality/#exactly) |  | equality |  |
| [Fail](testing/#fail) |  | testing |  |
| [FailNow](testing/#failnow) |  | testing |  |
| [FileEmpty](file/#fileempty) | [FileNotEmpty](file/#filenotempty) | file |  |
| [FileExists](file/#fileexists) | [FileNotExists](file/#filenotexists) | file |  |
| [Greater](comparison/#greater) | [LessOrEqual](comparison/#lessorequal) | comparison |  |
| [GreaterOrEqual](comparison/#greaterorequal) | [Less](comparison/#less) | comparison |  |
| [GreaterOrEqualT[Orderable Ordered]](comparison/#greaterorequaltorderable-ordered) {{% icon icon="star" color=orange %}} | [LessT](comparison/#lesstorderable-ordered) | comparison |  |
| [GreaterT[Orderable Ordered]](comparison/#greatertorderable-ordered) {{% icon icon="star" color=orange %}} | [LessOrEqualT](comparison/#lessorequaltorderable-ordered) | comparison |  |
| [HTTPBody](http/#httpbody) |  | http | helper |
| [HTTPBodyContains](http/#httpbodycontains) | [HTTPBodyNotContains](http/#httpbodynotcontains) | http |  |
| [HTTPError](http/#httperror) |  | http |  |
| [HTTPRedirect](http/#httpredirect) |  | http |  |
| [HTTPStatusCode](http/#httpstatuscode) |  | http |  |
| [HTTPSuccess](http/#httpsuccess) |  | http |  |
| [Implements](type/#implements) | [NotImplements](type/#notimplements) | type |  |
| [InDelta](number/#indelta) |  | number |  |
| [InDeltaMapValues](number/#indeltamapvalues) |  | number |  |
| [InDeltaSlice](number/#indeltaslice) |  | number |  |
| [InDeltaT[Number Measurable]](number/#indeltatnumber-measurable) {{% icon icon="star" color=orange %}} |  | number |  |
| [InEpsilon](number/#inepsilon) |  | number |  |
| [InEpsilonSlice](number/#inepsilonslice) |  | number |  |
| [InEpsilonT[Number Measurable]](number/#inepsilontnumber-measurable) {{% icon icon="star" color=orange %}} |  | number |  |
| [IsDecreasing](ordering/#isdecreasing) | [IsNonDecreasing](ordering/#isnondecreasing) | ordering |  |
| [IsDecreasingT[OrderedSlice ~[]E, E Ordered]](ordering/#isdecreasingtorderedslice-e-e-ordered) {{% icon icon="star" color=orange %}} | [IsNonDecreasingT](ordering/#isnondecreasingtorderedslice-e-e-ordered) | ordering |  |
| [IsIncreasing](ordering/#isincreasing) | [IsNonIncreasing](ordering/#isnonincreasing) | ordering |  |
| [IsIncreasingT[OrderedSlice ~[]E, E Ordered]](ordering/#isincreasingtorderedslice-e-e-ordered) {{% icon icon="star" color=orange %}} | [IsNonIncreasingT](ordering/#isnonincreasingtorderedslice-e-e-ordered) | ordering |  |
| [IsOfTypeT[EType any]](type/#isoftypetetype-any) {{% icon icon="star" color=orange %}} | [IsNotOfTypeT](type/#isnotoftypetetype-any) | type |  |
| [IsType](type/#istype) | [IsNotType](type/#isnottype) | type |  |
| [JSONEq](json/#jsoneq) |  | json |  |
| [JSONEqBytes](json/#jsoneqbytes) |  | json |  |
| [JSONEqT[EDoc, ADoc Text]](json/#jsoneqtedoc-adoc-text) {{% icon icon="star" color=orange %}} |  | json |  |
| [JSONMarshalAsT[EDoc Text]](json/#jsonmarshalastedoc-text) {{% icon icon="star" color=orange %}} |  | json |  |
| [JSONUnmarshalAsT[Object any, ADoc Text]](json/#jsonunmarshalastobject-any-adoc-text) {{% icon icon="star" color=orange %}} |  | json |  |
| [Kind](type/#kind) | [NotKind](type/#notkind) | type |  |
| [Len](collection/#len) |  | collection |  |
| [MapContainsT[Map ~map[K]V, K comparable, V any]](collection/#mapcontainstmap-mapkv-k-comparable-v-any) {{% icon icon="star" color=orange %}} | [MapNotContainsT](collection/#mapnotcontainstmap-mapkv-k-comparable-v-any) | collection |  |
| [MapEqualT[K, V comparable]](collection/#mapequaltk-v-comparable) {{% icon icon="star" color=orange %}} | [MapNotEqualT](collection/#mapnotequaltk-v-comparable) | collection |  |
| [Nil](equality/#nil) | [NotNil](equality/#notnil) | equality |  |
| [NoFileDescriptorLeak](safety/#nofiledescriptorleak) |  | safety |  |
| [NoGoRoutineLeak](safety/#nogoroutineleak) |  | safety |  |
| [ObjectsAreEqual](common/#objectsareequal) |  | common | helper |
| [ObjectsAreEqualValues](common/#objectsareequalvalues) |  | common | helper |
| [Panics](panic/#panics) | [NotPanics](panic/#notpanics) | panic |  |
| [PanicsWithError](panic/#panicswitherror) |  | panic |  |
| [PanicsWithValue](panic/#panicswithvalue) |  | panic |  |
| [Positive](comparison/#positive) | [Negative](comparison/#negative) | comparison |  |
| [PositiveT[SignedNumber SignedNumeric]](comparison/#positivetsignednumber-signednumeric) {{% icon icon="star" color=orange %}} | [NegativeT](comparison/#negativetsignednumber-signednumeric) | comparison |  |
| [Regexp](string/#regexp) | [NotRegexp](string/#notregexp) | string |  |
| [RegexpT[Rex RegExp, ADoc Text]](string/#regexptrex-regexp-adoc-text) {{% icon icon="star" color=orange %}} | [NotRegexpT](string/#notregexptrex-regexp-adoc-text) | string |  |
| [Same](equality/#same) | [NotSame](equality/#notsame) | equality |  |
| [SameT[P any]](equality/#sametp-any) {{% icon icon="star" color=orange %}} | [NotSameT](equality/#notsametp-any) | equality |  |
| [SeqContainsT[E comparable]](collection/#seqcontainste-comparable) {{% icon icon="star" color=orange %}} | [SeqNotContainsT](collection/#seqnotcontainste-comparable) | collection |  |
| [SliceContainsT[Slice ~[]E, E comparable]](collection/#slicecontainstslice-e-e-comparable) {{% icon icon="star" color=orange %}} | [SliceNotContainsT](collection/#slicenotcontainstslice-e-e-comparable) | collection |  |
| [SliceEqualT[E comparable]](collection/#sliceequalte-comparable) {{% icon icon="star" color=orange %}} | [SliceNotEqualT](collection/#slicenotequalte-comparable) | collection |  |
| [SliceSubsetT[Slice ~[]E, E comparable]](collection/#slicesubsettslice-e-e-comparable) {{% icon icon="star" color=orange %}} | [SliceNotSubsetT](collection/#slicenotsubsettslice-e-e-comparable) | collection |  |
| [SortedT[OrderedSlice ~[]E, E Ordered]](ordering/#sortedtorderedslice-e-e-ordered) {{% icon icon="star" color=orange %}} | [NotSortedT](ordering/#notsortedtorderedslice-e-e-ordered) | ordering |  |
| [StringContainsT[ADoc, EDoc Text]](collection/#stringcontainstadoc-edoc-text) {{% icon icon="star" color=orange %}} | [StringNotContainsT](collection/#stringnotcontainstadoc-edoc-text) | collection |  |
| [Subset](collection/#subset) | [NotSubset](collection/#notsubset) | collection |  |
| [True](boolean/#true) | [False](boolean/#false) | boolean |  |
| [TrueT[B Boolean]](boolean/#truetb-boolean) {{% icon icon="star" color=orange %}} | [FalseT](boolean/#falsetb-boolean) | boolean |  |
| [WithinDuration](time/#withinduration) |  | time |  |
| [WithinRange](time/#withinrange) |  | time |  |
| [YAMLEq](yaml/#yamleq) |  | yaml |  |
| [YAMLEqBytes](yaml/#yamleqbytes) |  | yaml |  |
| [YAMLEqT[EDoc, ADoc Text]](yaml/#yamleqtedoc-adoc-text) {{% icon icon="star" color=orange %}} |  | yaml |  |
| [YAMLMarshalAsT[EDoc Text]](yaml/#yamlmarshalastedoc-text) {{% icon icon="star" color=orange %}} |  | yaml |  |
| [YAMLUnmarshalAsT[Object any, ADoc Text]](yaml/#yamlunmarshalastobject-any-adoc-text) {{% icon icon="star" color=orange %}} |  | yaml |  |
| [Zero](type/#zero) | [NotZero](type/#notzero) | type |  |

