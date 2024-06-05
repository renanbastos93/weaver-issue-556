# weaver-issue-556
Simulate issue opened

panic stack trace
```sh
weaverdemo git:(main) ✗ go run .                                                    
╭───────────────────────────────────────────────────╮
│ app        : weaverdemo                           │
│ deployment : 5535fd76-3c62-43ad-af45-44c2b9e44f2a │
╰───────────────────────────────────────────────────╯
runtime: goroutine stack exceeds 1000000000-byte limit
runtime: sp=0x14020612350 stack=[0x14020612000, 0x14040612000]
fatal error: stack overflow

runtime stack:
runtime.throw({0x101160893?, 0x0?})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/panic.go:1023 +0x40 fp=0x1754a2d50 sp=0x1754a2d20 pc=0x100a1f9b0
runtime.newstack()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/stack.go:1103 +0x460 fp=0x1754a2f00 sp=0x1754a2d50 pc=0x100a3d040
runtime.morestack()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:341 +0x70 fp=0x1754a2f00 sp=0x1754a2f00 pc=0x100a56090

goroutine 1 gp=0x140000021c0 m=13 mp=0x14000600908 [running]:
runtime.mapaccess1_fast32(0x101347440, 0x0, 0x136d40)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/map_fast32.go:13 +0x184 fp=0x14020612350 sp=0x14020612350 pc=0x1009f8084
runtime.resolveTypeOff(0x0?, 0x136d40)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/type.go:170 +0x90 fp=0x140206123b0 sp=0x14020612350 pc=0x100a4c8e0
runtime.rtype.typeOff(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/type.go:182
reflect.resolveTypeOff(0x14020612408?, 0x9f44b8?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/runtime1.go:618 +0x1c fp=0x140206123d0 sp=0x140206123b0 pc=0x100a533bc
reflect.(*rtype).typeOff(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/reflect/type.go:537
reflect.(*rtype).ptrTo(0x10?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/reflect/type.go:1194 +0xa8 fp=0x14020612460 sp=0x140206123d0 pc=0x100aa5ec8
reflect.ptrTo(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/reflect/type.go:1236
reflect.New({0x101444890?, 0x1013e0f40})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/reflect/value.go:3300 +0x48 fp=0x14020612490 sp=0x14020612460 pc=0x100ab8548
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:281 +0x74 fp=0x14020612550 sp=0x14020612490 pc=0x101121f94
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117a7fe, 0x2e})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x140206125b0 sp=0x14020612550 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x140206125f0 sp=0x140206125b0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013d2a60, 0x14010bd1a40}, 0x14010bd81f8)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x140206127d0 sp=0x140206125f0 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e6e0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14020612890 sp=0x140206127d0 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101363ce0}, {0x10117b2f4, 0x2f})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x140206128f0 sp=0x14020612890 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101363ce0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14020612930 sp=0x140206128f0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013e6d40, 0x14010bd1a00}, 0x14010bd81e0)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14020612b10 sp=0x14020612930 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14020612bd0 sp=0x14020612b10 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117a7fe, 0x2e})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14020612c30 sp=0x14020612bd0 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14020612c70 sp=0x14020612c30 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013d2a60, 0x14010bd19c0}, 0x14010bd81c8)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14020612e50 sp=0x14020612c70 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e6e0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14020612f10 sp=0x14020612e50 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101363ce0}, {0x10117b2f4, 0x2f})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14020612f70 sp=0x14020612f10 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101363ce0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14020612fb0 sp=0x14020612f70 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013e6d40, 0x14010bd1980}, 0x14010bd81b0)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14020613190 sp=0x14020612fb0 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14020613250 sp=0x14020613190 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117a7fe, 0x2e})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x140206132b0 sp=0x14020613250 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x140206132f0 sp=0x140206132b0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013d2a60, 0x14010bd1940}, 0x14010bd8198)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x140206134d0 sp=0x140206132f0 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e6e0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14020613590 sp=0x140206134d0 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101363ce0}, {0x10117b2f4, 0x2f})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x140206135f0 sp=0x14020613590 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101363ce0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14020613630 sp=0x140206135f0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013e6d40, 0x14010bd1900}, 0x14010bd8180)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14020613810 sp=0x14020613630 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x140206138d0 sp=0x14020613810 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117a7fe, 0x2e})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14020613930 sp=0x140206138d0 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14020613970 sp=0x14020613930 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013d2a60, 0x14010bd18c0}, 0x14010bd8168)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14020613b50 sp=0x14020613970 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e6e0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14020613c10 sp=0x14020613b50 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101363ce0}, {0x10117b2f4, 0x2f})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14020613c70 sp=0x14020613c10 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101363ce0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14020613cb0 sp=0x14020613c70 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013e6d40, 0x14010bd1880}, 0x14010bd8150)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14020613e90 sp=0x14020613cb0 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14020613f50 sp=0x14020613e90 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117a7fe, 0x2e})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14020613fb0 sp=0x14020613f50 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14020613ff0 sp=0x14020613fb0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013d2a60, 0x14010bd1840}, 0x14010bd8138)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x140206141d0 sp=0x14020613ff0 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e6e0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14020614290 sp=0x140206141d0 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101363ce0}, {0x10117b2f4, 0x2f})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x140206142f0 sp=0x14020614290 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101363ce0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14020614330 sp=0x140206142f0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013e6d40, 0x14010bd1800}, 0x14010bd8120)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14020614510 sp=0x14020614330 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x140206145d0 sp=0x14020614510 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117a7fe, 0x2e})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14020614630 sp=0x140206145d0 pc=0x101121dc8
...2581016 frames elided...
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x1404060f9f0 sp=0x1404060f9b0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013d2a60, 0x140003ce280}, 0x1400000e108)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x1404060fbd0 sp=0x1404060f9f0 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e6e0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x1404060fc90 sp=0x1404060fbd0 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101363ce0}, {0x10117b2f4, 0x2f})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x1404060fcf0 sp=0x1404060fc90 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101363ce0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x1404060fd30 sp=0x1404060fcf0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013e6d40, 0x140003ce240}, 0x1400000e0f0)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x1404060ff10 sp=0x1404060fd30 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x1404060ffd0 sp=0x1404060ff10 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117a7fe, 0x2e})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14040610030 sp=0x1404060ffd0 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14040610070 sp=0x14040610030 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013d2a60, 0x140003ce200}, 0x1400000e0d8)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14040610250 sp=0x14040610070 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e6e0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14040610310 sp=0x14040610250 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101363ce0}, {0x10117b2f4, 0x2f})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14040610370 sp=0x14040610310 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101363ce0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x140406103b0 sp=0x14040610370 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013e6d40, 0x140003ce1c0}, 0x1400000e0c0)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14040610590 sp=0x140406103b0 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14040610650 sp=0x14040610590 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117a7fe, 0x2e})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x140406106b0 sp=0x14040610650 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x140406106f0 sp=0x140406106b0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013d2a60, 0x140003ce180}, 0x1400000e0a8)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x140406108d0 sp=0x140406106f0 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e6e0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14040610990 sp=0x140406108d0 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101363ce0}, {0x10117b2f4, 0x2f})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x140406109f0 sp=0x14040610990 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101363ce0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14040610a30 sp=0x140406109f0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013e6d40, 0x140003ce140}, 0x1400000e090)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14040610c10 sp=0x14040610a30 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14040610cd0 sp=0x14040610c10 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117a7fe, 0x2e})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14040610d30 sp=0x14040610cd0 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14040610d70 sp=0x14040610d30 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013d2a60, 0x140003ce100}, 0x1400000e078)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14040610f50 sp=0x14040610d70 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e6e0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14040611010 sp=0x14040610f50 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101363ce0}, {0x10117b2f4, 0x2f})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14040611070 sp=0x14040611010 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101363ce0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x140406110b0 sp=0x14040611070 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013e6d40, 0x140003ce0c0}, 0x1400000e060)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14040611290 sp=0x140406110b0 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14040611350 sp=0x14040611290 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117a7fe, 0x2e})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x140406113b0 sp=0x14040611350 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x140406113f0 sp=0x140406113b0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013d2a60, 0x140003ce080}, 0x1400000e048)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x140406115d0 sp=0x140406113f0 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e6e0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14040611690 sp=0x140406115d0 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101363ce0}, {0x10117b2f4, 0x2f})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x140406116f0 sp=0x14040611690 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101363ce0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14040611730 sp=0x140406116f0 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013e6d40, 0x140003ce040}, 0x1400000e030)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14040611910 sp=0x14040611730 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e5a0)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x140406119d0 sp=0x14040611910 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getIntf(0x140001623c0, {0x101444890, 0x101354120}, {0x10117335b, 0x24})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:246 +0xa8 fp=0x14040611a30 sp=0x140406119d0 pc=0x101121dc8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get.func1({0x101444890?, 0x101354120?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:303 +0x40 fp=0x14040611a70 sp=0x14040611a30 pc=0x101122570
github.com/ServiceWeaver/weaver.fillRefs({0x1013992a0, 0x140003ce000}, 0x1400000e018)
        /Users/unknow/git/folks/weaver/fill.go:107 +0x244 fp=0x14040611c50 sp=0x14040611a70 pc=0x10112af44
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).get(0x140001623c0, 0x1400019e820)
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:302 +0x2d8 fp=0x14040611d10 sp=0x14040611c50 pc=0x1011221f8
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).getImpl(0x140001623c0, {0x101444890, 0x1013d12c0})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:262 +0x9c fp=0x14040611d60 sp=0x14040611d10 pc=0x101121eec
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).GetImpl(0x101438b60?, {0x101444890?, 0x1013d12c0?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:234 +0xd8 fp=0x14040611dd0 sp=0x14040611d60 pc=0x101121bf8
github.com/ServiceWeaver/weaver.runLocal[...]({0x101438b60, 0x101a6df60}, 0x10142af50)
        /Users/unknow/git/folks/weaver/weaver.go:176 +0x1a0 fp=0x14040611eb0 sp=0x14040611dd0 pc=0x101153ee0
github.com/ServiceWeaver/weaver.Run[...]({0x101438b60, 0x101a6df60}, 0x10142af50)
        /Users/unknow/git/folks/weaver/weaver.go:143 +0xa8 fp=0x14040611f00 sp=0x14040611eb0 pc=0x101153a18
main.main()
        /Users/unknow/git/folks/weaverdemo/main.go:34 +0x3c fp=0x14040611f40 sp=0x14040611f00 pc=0x10114d90c
runtime.main()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:271 +0x28c fp=0x14040611fd0 sp=0x14040611f40 pc=0x100a222ec
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14040611fd0 sp=0x14040611fd0 pc=0x100a58464

goroutine 18 gp=0x14000104380 m=nil [force gc (idle)]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000072790 sp=0x14000072770 pc=0x100a22718
runtime.goparkunlock(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:408
runtime.forcegchelper()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:326 +0xb8 fp=0x140000727d0 sp=0x14000072790 pc=0x100a225a8
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140000727d0 sp=0x140000727d0 pc=0x100a58464
created by runtime.init.6 in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:314 +0x24

goroutine 19 gp=0x14000104540 m=nil [GC sweep wait]:
runtime.gopark(0x1?, 0x0?, 0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000072f60 sp=0x14000072f40 pc=0x100a22718
runtime.goparkunlock(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:408
runtime.bgsweep(0x14000112000)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgcsweep.go:318 +0x108 fp=0x14000072fb0 sp=0x14000072f60 pc=0x100a0d468
runtime.gcenable.gowrap1()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:203 +0x28 fp=0x14000072fd0 sp=0x14000072fb0 pc=0x100a01608
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14000072fd0 sp=0x14000072fd0 pc=0x100a58464
created by runtime.gcenable in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:203 +0x6c

goroutine 20 gp=0x14000104700 m=nil [GC scavenge wait]:
runtime.gopark(0x10000?, 0x10129ec28?, 0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000073760 sp=0x14000073740 pc=0x100a22718
runtime.goparkunlock(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:408
runtime.(*scavengerState).park(0x101a06940)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgcscavenge.go:425 +0x5c fp=0x14000073790 sp=0x14000073760 pc=0x100a0adec
runtime.bgscavenge(0x14000112000)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgcscavenge.go:658 +0xac fp=0x140000737b0 sp=0x14000073790 pc=0x100a0b3ac
runtime.gcenable.gowrap2()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:204 +0x28 fp=0x140000737d0 sp=0x140000737b0 pc=0x100a015a8
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140000737d0 sp=0x140000737d0 pc=0x100a58464
created by runtime.gcenable in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:204 +0xac

goroutine 34 gp=0x14000188380 m=nil [finalizer wait]:
runtime.gopark(0x140000765b8?, 0x100a52d08?, 0x8?, 0x0?, 0x101319ce0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000076580 sp=0x14000076560 pc=0x100a22718
runtime.runfinq()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mfinal.go:194 +0x108 fp=0x140000767d0 sp=0x14000076580 pc=0x100a006d8
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140000767d0 sp=0x140000767d0 pc=0x100a58464
created by runtime.createfing in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mfinal.go:164 +0x80

goroutine 35 gp=0x14000189880 m=nil [GC worker (idle)]:
runtime.gopark(0x101a709c0?, 0x3?, 0x2b?, 0xf5?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140004ea730 sp=0x140004ea710 pc=0x100a22718
runtime.gcBgMarkWorker()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1310 +0xd8 fp=0x140004ea7d0 sp=0x140004ea730 pc=0x100a03708
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140004ea7d0 sp=0x140004ea7d0 pc=0x100a58464
created by runtime.gcBgMarkStartWorkers in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1234 +0x28

goroutine 36 gp=0x14000189a40 m=nil [GC worker (idle)]:
runtime.gopark(0x101a709c0?, 0x3?, 0x5e?, 0x18?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140004eaf30 sp=0x140004eaf10 pc=0x100a22718
runtime.gcBgMarkWorker()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1310 +0xd8 fp=0x140004eafd0 sp=0x140004eaf30 pc=0x100a03708
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140004eafd0 sp=0x140004eafd0 pc=0x100a58464
created by runtime.gcBgMarkStartWorkers in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1234 +0x28

goroutine 2 gp=0x14000003340 m=nil [GC worker (idle)]:
runtime.gopark(0x101a709c0?, 0x3?, 0x9e?, 0x36?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000076f30 sp=0x14000076f10 pc=0x100a22718
runtime.gcBgMarkWorker()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1310 +0xd8 fp=0x14000076fd0 sp=0x14000076f30 pc=0x100a03708
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14000076fd0 sp=0x14000076fd0 pc=0x100a58464
created by runtime.gcBgMarkStartWorkers in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1234 +0x28

goroutine 37 gp=0x14000189c00 m=nil [GC worker (idle)]:
runtime.gopark(0x101a709c0?, 0x1?, 0xd5?, 0x2?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140004eb730 sp=0x140004eb710 pc=0x100a22718
runtime.gcBgMarkWorker()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1310 +0xd8 fp=0x140004eb7d0 sp=0x140004eb730 pc=0x100a03708
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140004eb7d0 sp=0x140004eb7d0 pc=0x100a58464
created by runtime.gcBgMarkStartWorkers in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1234 +0x28

goroutine 21 gp=0x140001048c0 m=nil [GC worker (idle)]:
runtime.gopark(0x196ffc466b019?, 0x0?, 0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000073f30 sp=0x14000073f10 pc=0x100a22718
runtime.gcBgMarkWorker()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1310 +0xd8 fp=0x14000073fd0 sp=0x14000073f30 pc=0x100a03708
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14000073fd0 sp=0x14000073fd0 pc=0x100a58464
created by runtime.gcBgMarkStartWorkers in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1234 +0x28

goroutine 38 gp=0x14000189dc0 m=nil [GC worker (idle)]:
runtime.gopark(0x101a709c0?, 0x3?, 0x57?, 0xb8?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140004ebf30 sp=0x140004ebf10 pc=0x100a22718
runtime.gcBgMarkWorker()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1310 +0xd8 fp=0x140004ebfd0 sp=0x140004ebf30 pc=0x100a03708
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140004ebfd0 sp=0x140004ebfd0 pc=0x100a58464
created by runtime.gcBgMarkStartWorkers in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1234 +0x28

goroutine 3 gp=0x14000003500 m=nil [GC worker (idle)]:
runtime.gopark(0x101a709c0?, 0x1?, 0xb6?, 0xb8?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000077730 sp=0x14000077710 pc=0x100a22718
runtime.gcBgMarkWorker()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1310 +0xd8 fp=0x140000777d0 sp=0x14000077730 pc=0x100a03708
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140000777d0 sp=0x140000777d0 pc=0x100a58464
created by runtime.gcBgMarkStartWorkers in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1234 +0x28

goroutine 39 gp=0x140004ee000 m=nil [GC worker (idle)]:
runtime.gopark(0x101a709c0?, 0x3?, 0x92?, 0xc7?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140004ec730 sp=0x140004ec710 pc=0x100a22718
runtime.gcBgMarkWorker()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1310 +0xd8 fp=0x140004ec7d0 sp=0x140004ec730 pc=0x100a03708
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140004ec7d0 sp=0x140004ec7d0 pc=0x100a58464
created by runtime.gcBgMarkStartWorkers in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/mgc.go:1234 +0x28

goroutine 22 gp=0x140004ef500 m=nil [select]:
runtime.gopark(0x14000074f78?, 0x2?, 0x68?, 0x4e?, 0x14000074f74?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x1400008be20 sp=0x1400008be00 pc=0x100a22718
runtime.selectgo(0x1400008bf78, 0x14000074f70, 0x0?, 0x0, 0x0?, 0x1)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/select.go:327 +0x614 fp=0x1400008bf30 sp=0x1400008be20 pc=0x100a35a84
database/sql.(*DB).connectionOpener(0x1400029add0, {0x101438d20, 0x140000c1bd0})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/database/sql/sql.go:1246 +0x80 fp=0x1400008bfa0 sp=0x1400008bf30 pc=0x100f10210
database/sql.OpenDB.gowrap1()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/database/sql/sql.go:824 +0x30 fp=0x1400008bfd0 sp=0x1400008bfa0 pc=0x100f0e410
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400008bfd0 sp=0x1400008bfd0 pc=0x100a58464
created by database/sql.OpenDB in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/database/sql/sql.go:824 +0x140

goroutine 23 gp=0x140004ef6c0 m=nil [select]:
runtime.gopark(0x14000075720?, 0x3?, 0x8?, 0x81?, 0x140000756e2?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000075570 sp=0x14000075550 pc=0x100a22718
runtime.selectgo(0x14000075720, 0x140000756dc, 0x0?, 0x0, 0x0?, 0x1)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/select.go:327 +0x614 fp=0x14000075680 sp=0x14000075570 pc=0x100a35a84
go.opentelemetry.io/otel/sdk/trace.(*batchSpanProcessor).processQueue(0x1400019ebe0)
        /Users/unknow/go/pkg/mod/go.opentelemetry.io/otel/sdk@v1.19.0/trace/batch_span_processor.go:312 +0xd4 fp=0x14000075790 sp=0x14000075680 pc=0x100eb5044
go.opentelemetry.io/otel/sdk/trace.NewBatchSpanProcessor.func1()
        /Users/unknow/go/pkg/mod/go.opentelemetry.io/otel/sdk@v1.19.0/trace/batch_span_processor.go:128 +0x58 fp=0x140000757d0 sp=0x14000075790 pc=0x100eb4398
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140000757d0 sp=0x140000757d0 pc=0x100a58464
created by go.opentelemetry.io/otel/sdk/trace.NewBatchSpanProcessor in goroutine 1
        /Users/unknow/go/pkg/mod/go.opentelemetry.io/otel/sdk@v1.19.0/trace/batch_span_processor.go:126 +0x2b8

goroutine 24 gp=0x140004ef880 m=nil [select, locked to thread]:
runtime.gopark(0x14000075fa0?, 0x2?, 0xa8?, 0x5e?, 0x14000075f90?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000075e40 sp=0x14000075e20 pc=0x100a22718
runtime.selectgo(0x14000075fa0, 0x14000075f8c, 0x0?, 0x0, 0x0?, 0x1)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/select.go:327 +0x614 fp=0x14000075f50 sp=0x14000075e40 pc=0x100a35a84
runtime.ensureSigM.func1()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/signal_unix.go:1034 +0x168 fp=0x14000075fd0 sp=0x14000075f50 pc=0x100a4e548
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14000075fd0 sp=0x14000075fd0 pc=0x100a58464
created by runtime.ensureSigM in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/signal_unix.go:1017 +0xd8

goroutine 40 gp=0x14000104a80 m=8 mp=0x14000186408 [syscall]:
runtime.sigNoteSleep(0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/os_darwin.go:132 +0x20 fp=0x140004e8790 sp=0x140004e8750 pc=0x100a1d030
os/signal.signal_recv()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/sigqueue.go:149 +0x2c fp=0x140004e87b0 sp=0x140004e8790 pc=0x100a53ecc
os/signal.loop()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/os/signal/signal_unix.go:23 +0x1c fp=0x140004e87d0 sp=0x140004e87b0 pc=0x100dafa1c
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140004e87d0 sp=0x140004e87d0 pc=0x100a58464
created by os/signal.Notify.func1.1 in goroutine 1
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/os/signal/signal.go:151 +0x28

goroutine 41 gp=0x14000104c40 m=nil [chan receive]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140004e8df0 sp=0x140004e8dd0 pc=0x100a22718
runtime.chanrecv(0x140002b6300, 0x0, 0x1)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/chan.go:583 +0x404 fp=0x140004e8e70 sp=0x140004e8df0 pc=0x1009ed724
runtime.chanrecv1(0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/chan.go:442 +0x14 fp=0x140004e8ea0 sp=0x140004e8e70 pc=0x1009ed314
github.com/ServiceWeaver/weaver/internal/weaver.NewSingleWeavelet.func1()
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:159 +0x4c fp=0x140004e8fd0 sp=0x140004e8ea0 pc=0x10112101c
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140004e8fd0 sp=0x140004e8fd0 pc=0x100a58464
created by github.com/ServiceWeaver/weaver/internal/weaver.NewSingleWeavelet in goroutine 1
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:158 +0x798

goroutine 42 gp=0x14000104e00 m=nil [select]:
runtime.gopark(0x140002e3468?, 0x6?, 0x68?, 0x31?, 0x140002e328c?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140002e3130 sp=0x140002e3110 pc=0x100a22718
runtime.selectgo(0x140002e3468, 0x140002e3280, 0x101161217?, 0x0, 0x140002b66c0?, 0x1)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/select.go:327 +0x614 fp=0x140002e3240 sp=0x140002e3130 pc=0x100a35a84
net/http.(*persistConn).roundTrip(0x14000000360, 0x14000279280)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/transport.go:2675 +0x7c8 fp=0x140002e34d0 sp=0x140002e3240 pc=0x100d69db8
net/http.(*Transport).roundTrip(0x1019f99a0, 0x14000000120)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/transport.go:608 +0x68c fp=0x140002e3700 sp=0x140002e34d0 pc=0x100d5ebec
net/http.(*Transport).RoundTrip(0x0?, 0x101433120?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/roundtrip.go:17 +0x1c fp=0x140002e3720 sp=0x140002e3700 pc=0x100d44ffc
net/http.send(0x14000000120, {0x101433120, 0x1019f99a0}, {0x100d02cc0?, 0x8?, 0x0?})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/client.go:259 +0x4ac fp=0x140002e3900 sp=0x140002e3720 pc=0x100d015fc
net/http.(*Client).send(0x101a05f80, 0x14000000120, {0x101a709c0?, 0x101af85e0?, 0x0?})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/client.go:180 +0x9c fp=0x140002e3980 sp=0x140002e3900 pc=0x100d00fec
net/http.(*Client).do(0x101a05f80, 0x14000000120)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/client.go:724 +0x6b4 fp=0x140002e3b70 sp=0x140002e3980 pc=0x100d02d14
net/http.(*Client).Do(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/client.go:590
github.com/ServiceWeaver/weaver/runtime/protomsg.Call({0x101438b60, 0x101a6df60}, {0x101a05f80, {0x0, 0x0}, {0x140001a85a0, 0x16}, {0x10116bc49, 0x1b}, {0x0, ...}, ...})
        /Users/unknow/git/folks/weaver/runtime/protomsg/call.go:81 +0x248 fp=0x140002e3cb0 sp=0x140002e3b70 pc=0x100dd2228
github.com/ServiceWeaver/weaver/internal/status.(*Client).Status(0x140002e3ef0, {0x101438b60, 0x101a6df60})
        /Users/unknow/git/folks/weaver/internal/status/client.go:41 +0xe0 fp=0x140002e3d90 sp=0x140002e3cb0 pc=0x1010ee010
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).ServeStatus(0x140001623c0, {0x101438b60, 0x101a6df60})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:407 +0x418 fp=0x140002e3f80 sp=0x140002e3d90 pc=0x101122bc8
github.com/ServiceWeaver/weaver.runLocal[...].func1()
        /Users/unknow/git/folks/weaver/weaver.go:171 +0x28 fp=0x140002e3fd0 sp=0x140002e3f80 pc=0x101152ec8
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140002e3fd0 sp=0x140002e3fd0 pc=0x100a58464
created by github.com/ServiceWeaver/weaver.runLocal[...] in goroutine 1
        /Users/unknow/git/folks/weaver/weaver.go:170 +0x184

goroutine 25 gp=0x140004efa40 m=nil [select]:
runtime.gopark(0x1400008ef10?, 0x2?, 0xe8?, 0x7d?, 0x1400008eeec?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x1400008ed90 sp=0x1400008ed70 pc=0x100a22718
runtime.selectgo(0x1400008ef10, 0x1400008eee8, 0x0?, 0x0, 0x0?, 0x1)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/select.go:327 +0x614 fp=0x1400008eea0 sp=0x1400008ed90 pc=0x100a35a84
github.com/ServiceWeaver/weaver/internal/metrics.(*StatsProcessor).CollectMetrics(0x14000299a40, {0x101438b60, 0x101a6df60}, 0x10142b388)
        /Users/unknow/git/folks/weaver/internal/metrics/stats.go:126 +0xe0 fp=0x1400008ef50 sp=0x1400008eea0 pc=0x100ed2b00
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).ServeStatus.func1()
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:381 +0x3c fp=0x1400008efd0 sp=0x1400008ef50 pc=0x101122f3c
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400008efd0 sp=0x1400008efd0 pc=0x100a58464
created by github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).ServeStatus in goroutine 42
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:380 +0x1e4

goroutine 28 gp=0x140004efc00 m=nil [select]:
runtime.gopark(0x14000078740?, 0x2?, 0x18?, 0x86?, 0x14000078724?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140000785d0 sp=0x140000785b0 pc=0x100a22718
runtime.selectgo(0x14000078740, 0x14000078720, 0x0?, 0x0, 0x0?, 0x1)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/select.go:327 +0x614 fp=0x140000786e0 sp=0x140000785d0 pc=0x100a35a84
github.com/ServiceWeaver/weaver/internal/weaver.serveHTTP({0x101438b60, 0x101a6df60}, {0x101437c70, 0x14000510360}, {0x101432c40, 0x1400029b1e0})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:538 +0x15c fp=0x14000078770 sp=0x140000786e0 pc=0x101123e7c
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).ServeStatus.func2()
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:401 +0x40 fp=0x140000787d0 sp=0x14000078770 pc=0x101122ed0
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140000787d0 sp=0x140000787d0 pc=0x100a58464
created by github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).ServeStatus in goroutine 42
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:400 +0x34c

goroutine 4 gp=0x14000504380 m=nil [IO wait]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x1009fe200?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x1400008fbc0 sp=0x1400008fba0 pc=0x100a22718
runtime.netpollblock(0x1400008fc58?, 0xacc314?, 0x1?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/netpoll.go:573 +0x158 fp=0x1400008fc00 sp=0x1400008fbc0 pc=0x100a1bf68
internal/poll.runtime_pollWait(0x101da5f00, 0x72)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/netpoll.go:345 +0xa0 fp=0x1400008fc30 sp=0x1400008fc00 pc=0x100a51ab0
internal/poll.(*pollDesc).wait(0x1400024ba00?, 0x1009fd11c?, 0x0)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/internal/poll/fd_poll_runtime.go:84 +0x28 fp=0x1400008fc60 sp=0x1400008fc30 pc=0x100ac7978
internal/poll.(*pollDesc).waitRead(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/internal/poll/fd_poll_runtime.go:89
internal/poll.(*FD).Accept(0x1400024ba00)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/internal/poll/fd_unix.go:611 +0x250 fp=0x1400008fd10 sp=0x1400008fc60 pc=0x100acc400
net.(*netFD).accept(0x1400024ba00)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/fd_unix.go:172 +0x28 fp=0x1400008fdd0 sp=0x1400008fd10 pc=0x100c57058
net.(*TCPListener).accept(0x14000510360)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/tcpsock_posix.go:159 +0x28 fp=0x1400008fe00 sp=0x1400008fdd0 pc=0x100c6a5f8
net.(*TCPListener).Accept(0x14000510360)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/tcpsock.go:327 +0x2c fp=0x1400008fe40 sp=0x1400008fe00 pc=0x100c6986c
net/http.(*onceCloseListener).Accept(0x14000310000?)
        <autogenerated>:1 +0x30 fp=0x1400008fe60 sp=0x1400008fe40 pc=0x100d75a40
net/http.(*Server).Serve(0x14000308000, {0x101437c70, 0x14000510360})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/server.go:3255 +0x2a8 fp=0x1400008ff90 sp=0x1400008fe60 pc=0x100d535f8
github.com/ServiceWeaver/weaver/internal/weaver.serveHTTP.func1()
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:537 +0x30 fp=0x1400008ffd0 sp=0x1400008ff90 pc=0x101123f20
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400008ffd0 sp=0x1400008ffd0 pc=0x100a58464
created by github.com/ServiceWeaver/weaver/internal/weaver.serveHTTP in goroutine 28
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:537 +0x110

goroutine 31 gp=0x140000036c0 m=nil [IO wait]:
runtime.gopark(0xffffffffffffffff?, 0xffffffffffffffff?, 0x23?, 0x0?, 0x100a6e4c0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x1400008da50 sp=0x1400008da30 pc=0x100a22718
runtime.netpollblock(0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/netpoll.go:573 +0x158 fp=0x1400008da90 sp=0x1400008da50 pc=0x100a1bf68
internal/poll.runtime_pollWait(0x101da5e08, 0x72)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/netpoll.go:345 +0xa0 fp=0x1400008dac0 sp=0x1400008da90 pc=0x100a51ab0
internal/poll.(*pollDesc).wait(0x1400024bb00?, 0x1400029d000?, 0x0)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/internal/poll/fd_poll_runtime.go:84 +0x28 fp=0x1400008daf0 sp=0x1400008dac0 pc=0x100ac7978
internal/poll.(*pollDesc).waitRead(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/internal/poll/fd_poll_runtime.go:89
internal/poll.(*FD).Read(0x1400024bb00, {0x1400029d000, 0x1000, 0x1000})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/internal/poll/fd_unix.go:164 +0x200 fp=0x1400008db90 sp=0x1400008daf0 pc=0x100ac8cc0
net.(*netFD).Read(0x1400024bb00, {0x1400029d000?, 0x0?, 0x100a42a2c?})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/fd_posix.go:55 +0x28 fp=0x1400008dbe0 sp=0x1400008db90 pc=0x100c55478
net.(*conn).Read(0x14000196d30, {0x1400029d000?, 0x140004e94d8?, 0x1009ec50c?})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/net.go:179 +0x34 fp=0x1400008dc30 sp=0x1400008dbe0 pc=0x100c61894
net.(*TCPConn).Read(0x140004efdc0?, {0x1400029d000?, 0x60?, 0x58?})
        <autogenerated>:1 +0x2c fp=0x1400008dc60 sp=0x1400008dc30 pc=0x100c7279c
net/http.(*persistConn).Read(0x14000000360, {0x1400029d000?, 0x140004e9518?, 0x100d676a0?})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/transport.go:1977 +0x50 fp=0x1400008dcc0 sp=0x1400008dc60 pc=0x100d667c0
bufio.(*Reader).fill(0x140002b6900)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/bufio/bufio.go:110 +0xf8 fp=0x1400008dd00 sp=0x1400008dcc0 pc=0x100b58508
bufio.(*Reader).Peek(0x140002b6900, 0x1)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/bufio/bufio.go:148 +0x60 fp=0x1400008dd20 sp=0x1400008dd00 pc=0x100b58670
net/http.(*persistConn).readLoop(0x14000000360)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/transport.go:2141 +0x158 fp=0x1400008dfb0 sp=0x1400008dd20 pc=0x100d67748
net/http.(*Transport).dialConn.gowrap2()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/transport.go:1799 +0x28 fp=0x1400008dfd0 sp=0x1400008dfb0 pc=0x100d65d58
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400008dfd0 sp=0x1400008dfd0 pc=0x100a58464
created by net/http.(*Transport).dialConn in goroutine 29
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/transport.go:1799 +0x1018

goroutine 5 gp=0x14000504540 m=nil [sync.Mutex.Lock]:
runtime.gopark(0x101a0f680?, 0x100a422d4?, 0x60?, 0xe3?, 0x140002df5c8?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140002df570 sp=0x140002df550 pc=0x100a22718
runtime.goparkunlock(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:408
runtime.semacquire1(0x1400016249c, 0x0, 0x3, 0x1, 0x15)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/sema.go:160 +0x208 fp=0x140002df5c0 sp=0x140002df570 pc=0x100a368c8
sync.runtime_SemacquireMutex(0x140002df628?, 0x34?, 0x140002df5f8?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/sema.go:77 +0x28 fp=0x140002df600 sp=0x140002df5c0 pc=0x100a53788
sync.(*Mutex).lockSlow(0x14000162498)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/sync/mutex.go:171 +0x174 fp=0x140002df650 sp=0x140002df600 pc=0x100a63784
sync.(*Mutex).Lock(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/sync/mutex.go:90
github.com/ServiceWeaver/weaver/internal/weaver.(*SingleWeavelet).Status(0x140001623c0, {0x1013b21e0?, 0x1013e8901?})
        /Users/unknow/git/folks/weaver/internal/weaver/singleweavelet.go:449 +0x98 fp=0x140002df870 sp=0x140002df650 pc=0x101123058
github.com/ServiceWeaver/weaver/internal/status.Server.Status-fm({0x101438d20?, 0x14000300190?})
        <autogenerated>:1 +0x44 fp=0x140002df8a0 sp=0x140002df870 pc=0x1010f9f34
github.com/ServiceWeaver/weaver/internal/status.RegisterServer.HandlerThunk[...].func2(0x14000620000)
        /Users/unknow/git/folks/weaver/runtime/protomsg/handler.go:123 +0x58 fp=0x140002df9b0 sp=0x140002df8a0 pc=0x1010f5fb8
github.com/ServiceWeaver/weaver/internal/status.RegisterServer.HandlerThunk[...].panicHandler.func5(0x14000312004?)
        /Users/unknow/git/folks/weaver/runtime/protomsg/handler.go:170 +0x6c fp=0x140002dfa10 sp=0x140002df9b0 pc=0x1010f5dac
github.com/ServiceWeaver/weaver/internal/status.RegisterServer.HandlerThunk[...].metricHandler.func6(0x14000620000)
        /Users/unknow/git/folks/weaver/runtime/protomsg/handler.go:185 +0x104 fp=0x140002dfac0 sp=0x140002dfa10 pc=0x1010f5c54
net/http.HandlerFunc.ServeHTTP(0x1400029b270?, {0x101438030?, 0x14000234540?}, 0x100d48da8?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/server.go:2166 +0x38 fp=0x140002dfaf0 sp=0x140002dfac0 pc=0x100d4fd68
net/http.(*ServeMux).ServeHTTP(0x0?, {0x101438030, 0x14000234540}, 0x14000620000)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/server.go:2683 +0x1a4 fp=0x140002dfb40 sp=0x140002dfaf0 pc=0x100d51944
net/http.serverHandler.ServeHTTP({0x14000612030?}, {0x101438030?, 0x14000234540?}, 0x6?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/server.go:3137 +0xbc fp=0x140002dfb70 sp=0x140002dfb40 pc=0x100d5324c
net/http.(*conn).serve(0x14000310000, {0x101438ce8, 0x1400030a120})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/server.go:2039 +0x508 fp=0x140002dffa0 sp=0x140002dfb70 pc=0x100d4e938
net/http.(*Server).Serve.gowrap3()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/server.go:3285 +0x30 fp=0x140002dffd0 sp=0x140002dffa0 pc=0x100d53990
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140002dffd0 sp=0x140002dffd0 pc=0x100a58464
created by net/http.(*Server).Serve in goroutine 4
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/server.go:3285 +0x3f0

goroutine 32 gp=0x14000003880 m=nil [select]:
runtime.gopark(0x140002ddf38?, 0x2?, 0xf8?, 0xdd?, 0x140002ddee4?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140002ddd90 sp=0x140002ddd70 pc=0x100a22718
runtime.selectgo(0x140002ddf38, 0x140002ddee0, 0x14000279340?, 0x0, 0x140002e6150?, 0x1)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/select.go:327 +0x614 fp=0x140002ddea0 sp=0x140002ddd90 pc=0x100a35a84
net/http.(*persistConn).writeLoop(0x14000000360)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/transport.go:2444 +0xa0 fp=0x140002ddfb0 sp=0x140002ddea0 pc=0x100d68ff0
net/http.(*Transport).dialConn.gowrap3()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/transport.go:1800 +0x28 fp=0x140002ddfd0 sp=0x140002ddfb0 pc=0x100d65cf8
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140002ddfd0 sp=0x140002ddfd0 pc=0x100a58464
created by net/http.(*Transport).dialConn in goroutine 29
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/transport.go:1800 +0x1060

goroutine 6 gp=0x14000504700 m=nil [IO wait]:
runtime.gopark(0xffffffffffffffff?, 0xffffffffffffffff?, 0x23?, 0x0?, 0x100a6e4c0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/proc.go:402 +0xc8 fp=0x140004e9550 sp=0x140004e9530 pc=0x100a22718
runtime.netpollblock(0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/netpoll.go:573 +0x158 fp=0x140004e9590 sp=0x140004e9550 pc=0x100a1bf68
internal/poll.runtime_pollWait(0x101da5d10, 0x72)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/netpoll.go:345 +0xa0 fp=0x140004e95c0 sp=0x140004e9590 pc=0x100a51ab0
internal/poll.(*pollDesc).wait(0x140000a8080?, 0x14000612041?, 0x0)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/internal/poll/fd_poll_runtime.go:84 +0x28 fp=0x140004e95f0 sp=0x140004e95c0 pc=0x100ac7978
internal/poll.(*pollDesc).waitRead(...)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/internal/poll/fd_poll_runtime.go:89
internal/poll.(*FD).Read(0x140000a8080, {0x14000612041, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/internal/poll/fd_unix.go:164 +0x200 fp=0x140004e9690 sp=0x140004e95f0 pc=0x100ac8cc0
net.(*netFD).Read(0x140000a8080, {0x14000612041?, 0x140004e9728?, 0x100c56e04?})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/fd_posix.go:55 +0x28 fp=0x140004e96e0 sp=0x140004e9690 pc=0x100c55478
net.(*conn).Read(0x14000306018, {0x14000612041?, 0x4?, 0x0?})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/net.go:179 +0x34 fp=0x140004e9730 sp=0x140004e96e0 pc=0x100c61894
net.(*TCPConn).Read(0x100d685b0?, {0x14000612041?, 0x140004e9618?, 0x100d686c0?})
        <autogenerated>:1 +0x2c fp=0x140004e9760 sp=0x140004e9730 pc=0x100c7279c
net/http.(*connReader).backgroundRead(0x14000612030)
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/server.go:681 +0x40 fp=0x140004e97b0 sp=0x140004e9760 pc=0x100d48f60
net/http.(*connReader).startBackgroundRead.gowrap2()
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/server.go:677 +0x28 fp=0x140004e97d0 sp=0x140004e97b0 pc=0x100d48e48
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140004e97d0 sp=0x140004e97d0 pc=0x100a58464
created by net/http.(*connReader).startBackgroundRead in goroutine 5
        /opt/homebrew/Cellar/go/1.22.3/libexec/src/net/http/server.go:677 +0xc8
exit status 2
```