// Code generated by serde. DO NOT EDIT

//go:build durable

package coroutine

import (
	ast "go/ast"
	unsafe "unsafe"
	unicode "unicode"
	types "go/types"
	token "go/token"
	time "time"
	syscall "syscall"
	syntax "regexp/syntax"
	sync "sync"
	strings "strings"
	strconv "strconv"
	sort "sort"
	slog "log/slog"
	scanner "go/scanner"
	runtime "runtime"
	regexp "regexp"
	reflect "reflect"
	parser "go/parser"
	os "os"
	log "log"
	io "io"
	fs "io/fs"
	exec "os/exec"
	doc "go/doc"
	crypto "crypto"
	constraint "go/build/constraint"
	constant "go/constant"
	comment "go/doc/comment"
	bytes "bytes"
	build "go/build"
	bufio "bufio"
	big "math/big"
	rand "math/rand"
	base64 "encoding/base64"
	json "encoding/json"
	atomic "sync/atomic"
	astutil "golang.org/x/tools/go/ast/astutil"
	serde "github.com/stealthrocket/coroutine/serde"
	semver "golang.org/x/mod/semver"
	packages "golang.org/x/tools/go/packages"
	objectpath "golang.org/x/tools/go/types/objectpath"
)

func init() {
	serde.RegisterType[**byte]()
	serde.RegisterType[*[100000]uintptr]()
	serde.RegisterType[*[1125899906842623]byte]()
	serde.RegisterType[*[131072]uint16]()
	serde.RegisterType[*[140737488355327]byte]()
	serde.RegisterType[*[16]byte]()
	serde.RegisterType[*[171]uint8]()
	serde.RegisterType[*[1]uintptr]()
	serde.RegisterType[*[268435456]uintptr]()
	serde.RegisterType[*[281474976710655]uint32]()
	serde.RegisterType[*[2]byte]()
	serde.RegisterType[*[2]float32]()
	serde.RegisterType[*[2]float64]()
	serde.RegisterType[*[2]int32]()
	serde.RegisterType[*[2]uint32]()
	serde.RegisterType[*[2]uintptr]()
	serde.RegisterType[*[32]rune]()
	serde.RegisterType[*[32]uintptr]()
	serde.RegisterType[*[4]byte]()
	serde.RegisterType[*[4]uint64]()
	serde.RegisterType[*[562949953421311]uint16]()
	serde.RegisterType[*[5]float64]()
	serde.RegisterType[*[65536]uintptr]()
	serde.RegisterType[*[6]float64]()
	serde.RegisterType[*[70368744177663]uint16]()
	serde.RegisterType[*[8]byte]()
	serde.RegisterType[*[8]uint8]()
	serde.RegisterType[*[][]byte]()
	serde.RegisterType[*[]byte]()
	serde.RegisterType[*[]rune]()
	serde.RegisterType[*[]string]()
	serde.RegisterType[*[]uint64]()
	serde.RegisterType[*bool]()
	serde.RegisterType[*byte]()
	serde.RegisterType[*complex128]()
	serde.RegisterType[*complex64]()
	serde.RegisterType[*float32]()
	serde.RegisterType[*float64]()
	serde.RegisterType[*int]()
	serde.RegisterType[*int16]()
	serde.RegisterType[*int32]()
	serde.RegisterType[*int64]()
	serde.RegisterType[*int8]()
	serde.RegisterType[*string]()
	serde.RegisterType[*uint]()
	serde.RegisterType[*uint16]()
	serde.RegisterType[*uint32]()
	serde.RegisterType[*uint64]()
	serde.RegisterType[*uint8]()
	serde.RegisterType[*uintptr]()
	serde.RegisterType[Frame]()
	serde.RegisterType[Stack]()
	serde.RegisterType[Storage]()
	serde.RegisterType[[0]byte]()
	serde.RegisterType[[0]int]()
	serde.RegisterType[[0]uint8]()
	serde.RegisterType[[0]uintptr]()
	serde.RegisterType[[1000]uintptr]()
	serde.RegisterType[[100]byte]()
	serde.RegisterType[[1024]bool]()
	serde.RegisterType[[1024]byte]()
	serde.RegisterType[[1024]uint8]()
	serde.RegisterType[[1048576]uint8]()
	serde.RegisterType[[104]byte]()
	serde.RegisterType[[108]byte]()
	serde.RegisterType[[108]int8]()
	serde.RegisterType[[10][]string]()
	serde.RegisterType[[10]byte]()
	serde.RegisterType[[10]float64]()
	serde.RegisterType[[10]string]()
	serde.RegisterType[[10]uint32]()
	serde.RegisterType[[11]float64]()
	serde.RegisterType[[11]string]()
	serde.RegisterType[[128][4]uint64]()
	serde.RegisterType[[128]bool]()
	serde.RegisterType[[128]byte]()
	serde.RegisterType[[128]float32]()
	serde.RegisterType[[128]uint16]()
	serde.RegisterType[[128]uint32]()
	serde.RegisterType[[128]uint64]()
	serde.RegisterType[[128]uintptr]()
	serde.RegisterType[[129]uint8]()
	serde.RegisterType[[12]float64]()
	serde.RegisterType[[131072]uintptr]()
	serde.RegisterType[[133]string]()
	serde.RegisterType[[13]int32]()
	serde.RegisterType[[147]uint16]()
	serde.RegisterType[[14]byte]()
	serde.RegisterType[[14]int8]()
	serde.RegisterType[[15]float64]()
	serde.RegisterType[[15]uint64]()
	serde.RegisterType[[16384]byte]()
	serde.RegisterType[[16384]uint8]()
	serde.RegisterType[[16]byte]()
	serde.RegisterType[[16]int64]()
	serde.RegisterType[[16]uint64]()
	serde.RegisterType[[16]uintptr]()
	serde.RegisterType[[17]byte]()
	serde.RegisterType[[17]string]()
	serde.RegisterType[[1]byte]()
	serde.RegisterType[[1]struct{}]()
	serde.RegisterType[[1]uint64]()
	serde.RegisterType[[1]uint8]()
	serde.RegisterType[[1]uintptr]()
	serde.RegisterType[[20]byte]()
	serde.RegisterType[[20]string]()
	serde.RegisterType[[20]uint64]()
	serde.RegisterType[[20]uint8]()
	serde.RegisterType[[21]byte]()
	serde.RegisterType[[21]string]()
	serde.RegisterType[[23]uint64]()
	serde.RegisterType[[249]uint8]()
	serde.RegisterType[[24]byte]()
	serde.RegisterType[[24]uint32]()
	serde.RegisterType[[252]uintptr]()
	serde.RegisterType[[253]uintptr]()
	serde.RegisterType[[256][]byte]()
	serde.RegisterType[[256]byte]()
	serde.RegisterType[[256]float32]()
	serde.RegisterType[[256]int]()
	serde.RegisterType[[256]int8]()
	serde.RegisterType[[256]uint32]()
	serde.RegisterType[[256]uint64]()
	serde.RegisterType[[256]uint8]()
	serde.RegisterType[[27]uint16]()
	serde.RegisterType[[28]uint64]()
	serde.RegisterType[[29]uint16]()
	serde.RegisterType[[2]byte]()
	serde.RegisterType[[2]int]()
	serde.RegisterType[[2]int32]()
	serde.RegisterType[[2]rune]()
	serde.RegisterType[[2]string]()
	serde.RegisterType[[2]uint16]()
	serde.RegisterType[[2]uint64]()
	serde.RegisterType[[2]uintptr]()
	serde.RegisterType[[32]byte]()
	serde.RegisterType[[32]float64]()
	serde.RegisterType[[32]string]()
	serde.RegisterType[[32]uint8]()
	serde.RegisterType[[32]uintptr]()
	serde.RegisterType[[33]byte]()
	serde.RegisterType[[33]float64]()
	serde.RegisterType[[3]byte]()
	serde.RegisterType[[3]float64]()
	serde.RegisterType[[3]int]()
	serde.RegisterType[[3]int64]()
	serde.RegisterType[[3]uint16]()
	serde.RegisterType[[3]uint32]()
	serde.RegisterType[[3]uint64]()
	serde.RegisterType[[4096]byte]()
	serde.RegisterType[[40]byte]()
	serde.RegisterType[[40]uint16]()
	serde.RegisterType[[44]byte]()
	serde.RegisterType[[4]byte]()
	serde.RegisterType[[4]float64]()
	serde.RegisterType[[4]int]()
	serde.RegisterType[[4]int64]()
	serde.RegisterType[[4]string]()
	serde.RegisterType[[4]uint16]()
	serde.RegisterType[[4]uint32]()
	serde.RegisterType[[4]uint64]()
	serde.RegisterType[[4]uint8]()
	serde.RegisterType[[4]uintptr]()
	serde.RegisterType[[50]byte]()
	serde.RegisterType[[50]uintptr]()
	serde.RegisterType[[512]byte]()
	serde.RegisterType[[512]uintptr]()
	serde.RegisterType[[51]uint16]()
	serde.RegisterType[[5]byte]()
	serde.RegisterType[[5]float64]()
	serde.RegisterType[[5]int]()
	serde.RegisterType[[5]uint]()
	serde.RegisterType[[5]uintptr]()
	serde.RegisterType[[607]int64]()
	serde.RegisterType[[61]struct{Size uint32; Mallocs uint64; Frees uint64}]()
	serde.RegisterType[[64488]byte]()
	serde.RegisterType[[64]byte]()
	serde.RegisterType[[64]uintptr]()
	serde.RegisterType[[65528]byte]()
	serde.RegisterType[[65]byte]()
	serde.RegisterType[[65]int8]()
	serde.RegisterType[[65]uint32]()
	serde.RegisterType[[65]uintptr]()
	serde.RegisterType[[67]uint16]()
	serde.RegisterType[[68]byte]()
	serde.RegisterType[[68]struct{Size uint32; Mallocs uint64; Frees uint64}]()
	serde.RegisterType[[68]uint16]()
	serde.RegisterType[[68]uint32]()
	serde.RegisterType[[68]uint64]()
	serde.RegisterType[[68]uint8]()
	serde.RegisterType[[696][2]uint64]()
	serde.RegisterType[[6]byte]()
	serde.RegisterType[[6]float64]()
	serde.RegisterType[[6]int]()
	serde.RegisterType[[6]int8]()
	serde.RegisterType[[6]uintptr]()
	serde.RegisterType[[72]byte]()
	serde.RegisterType[[768]byte]()
	serde.RegisterType[[7]float64]()
	serde.RegisterType[[7]string]()
	serde.RegisterType[[7]uint8]()
	serde.RegisterType[[800]byte]()
	serde.RegisterType[[8192]byte]()
	serde.RegisterType[[89]string]()
	serde.RegisterType[[8]byte]()
	serde.RegisterType[[8]float64]()
	serde.RegisterType[[8]string]()
	serde.RegisterType[[8]uint32]()
	serde.RegisterType[[8]uint64]()
	serde.RegisterType[[8]uint8]()
	serde.RegisterType[[96]byte]()
	serde.RegisterType[[96]int8]()
	serde.RegisterType[[9]string]()
	serde.RegisterType[[9]uintptr]()
	serde.RegisterType[[]*byte]()
	serde.RegisterType[[][2]int]()
	serde.RegisterType[[][2]string]()
	serde.RegisterType[[][2]uint16]()
	serde.RegisterType[[][][]byte]()
	serde.RegisterType[[][]byte]()
	serde.RegisterType[[][]int]()
	serde.RegisterType[[][]int32]()
	serde.RegisterType[[][]rune]()
	serde.RegisterType[[][]string]()
	serde.RegisterType[[]bool]()
	serde.RegisterType[[]byte]()
	serde.RegisterType[[]float32]()
	serde.RegisterType[[]float64]()
	serde.RegisterType[[]int]()
	serde.RegisterType[[]int16]()
	serde.RegisterType[[]int32]()
	serde.RegisterType[[]int64]()
	serde.RegisterType[[]int8]()
	serde.RegisterType[[]rune]()
	serde.RegisterType[[]string]()
	serde.RegisterType[[]struct{}]()
	serde.RegisterType[[]uint16]()
	serde.RegisterType[[]uint32]()
	serde.RegisterType[[]uint64]()
	serde.RegisterType[[]uint8]()
	serde.RegisterType[[]uintptr]()
	serde.RegisterType[ast.ArrayType]()
	serde.RegisterType[ast.AssignStmt]()
	serde.RegisterType[ast.BadDecl]()
	serde.RegisterType[ast.BadExpr]()
	serde.RegisterType[ast.BadStmt]()
	serde.RegisterType[ast.BasicLit]()
	serde.RegisterType[ast.BinaryExpr]()
	serde.RegisterType[ast.BlockStmt]()
	serde.RegisterType[ast.BranchStmt]()
	serde.RegisterType[ast.CallExpr]()
	serde.RegisterType[ast.CaseClause]()
	serde.RegisterType[ast.ChanDir]()
	serde.RegisterType[ast.ChanType]()
	serde.RegisterType[ast.CommClause]()
	serde.RegisterType[ast.Comment]()
	serde.RegisterType[ast.CommentGroup]()
	serde.RegisterType[ast.CompositeLit]()
	serde.RegisterType[ast.DeclStmt]()
	serde.RegisterType[ast.DeferStmt]()
	serde.RegisterType[ast.Ellipsis]()
	serde.RegisterType[ast.EmptyStmt]()
	serde.RegisterType[ast.ExprStmt]()
	serde.RegisterType[ast.Field]()
	serde.RegisterType[ast.FieldList]()
	serde.RegisterType[ast.File]()
	serde.RegisterType[ast.ForStmt]()
	serde.RegisterType[ast.FuncDecl]()
	serde.RegisterType[ast.FuncLit]()
	serde.RegisterType[ast.FuncType]()
	serde.RegisterType[ast.GenDecl]()
	serde.RegisterType[ast.GoStmt]()
	serde.RegisterType[ast.Ident]()
	serde.RegisterType[ast.IfStmt]()
	serde.RegisterType[ast.ImportSpec]()
	serde.RegisterType[ast.IncDecStmt]()
	serde.RegisterType[ast.IndexExpr]()
	serde.RegisterType[ast.IndexListExpr]()
	serde.RegisterType[ast.InterfaceType]()
	serde.RegisterType[ast.KeyValueExpr]()
	serde.RegisterType[ast.LabeledStmt]()
	serde.RegisterType[ast.MapType]()
	serde.RegisterType[ast.MergeMode]()
	serde.RegisterType[ast.ObjKind]()
	serde.RegisterType[ast.Object]()
	serde.RegisterType[ast.Package]()
	serde.RegisterType[ast.ParenExpr]()
	serde.RegisterType[ast.RangeStmt]()
	serde.RegisterType[ast.ReturnStmt]()
	serde.RegisterType[ast.Scope]()
	serde.RegisterType[ast.SelectStmt]()
	serde.RegisterType[ast.SelectorExpr]()
	serde.RegisterType[ast.SendStmt]()
	serde.RegisterType[ast.SliceExpr]()
	serde.RegisterType[ast.StarExpr]()
	serde.RegisterType[ast.StructType]()
	serde.RegisterType[ast.SwitchStmt]()
	serde.RegisterType[ast.TypeAssertExpr]()
	serde.RegisterType[ast.TypeSpec]()
	serde.RegisterType[ast.TypeSwitchStmt]()
	serde.RegisterType[ast.UnaryExpr]()
	serde.RegisterType[ast.ValueSpec]()
	serde.RegisterType[astutil.Cursor]()
	serde.RegisterType[atomic.Bool]()
	serde.RegisterType[atomic.Int32]()
	serde.RegisterType[atomic.Int64]()
	serde.RegisterType[atomic.Uint32]()
	serde.RegisterType[atomic.Uint64]()
	serde.RegisterType[atomic.Uintptr]()
	serde.RegisterType[atomic.Value]()
	serde.RegisterType[base64.CorruptInputError]()
	serde.RegisterType[base64.Encoding]()
	serde.RegisterType[big.Accuracy]()
	serde.RegisterType[big.ErrNaN]()
	serde.RegisterType[big.Float]()
	serde.RegisterType[big.Int]()
	serde.RegisterType[big.Rat]()
	serde.RegisterType[big.RoundingMode]()
	serde.RegisterType[big.Word]()
	serde.RegisterType[bool]()
	serde.RegisterType[bufio.ReadWriter]()
	serde.RegisterType[bufio.Reader]()
	serde.RegisterType[bufio.Scanner]()
	serde.RegisterType[bufio.Writer]()
	serde.RegisterType[build.Context]()
	serde.RegisterType[build.Directive]()
	serde.RegisterType[build.ImportMode]()
	serde.RegisterType[build.MultiplePackageError]()
	serde.RegisterType[build.NoGoError]()
	serde.RegisterType[build.Package]()
	serde.RegisterType[byte]()
	serde.RegisterType[bytes.Buffer]()
	serde.RegisterType[bytes.Reader]()
	serde.RegisterType[comment.Code]()
	serde.RegisterType[comment.Doc]()
	serde.RegisterType[comment.DocLink]()
	serde.RegisterType[comment.Heading]()
	serde.RegisterType[comment.Italic]()
	serde.RegisterType[comment.Link]()
	serde.RegisterType[comment.LinkDef]()
	serde.RegisterType[comment.List]()
	serde.RegisterType[comment.ListItem]()
	serde.RegisterType[comment.Paragraph]()
	serde.RegisterType[comment.Parser]()
	serde.RegisterType[comment.Plain]()
	serde.RegisterType[comment.Printer]()
	serde.RegisterType[complex128]()
	serde.RegisterType[complex64]()
	serde.RegisterType[constant.Kind]()
	serde.RegisterType[constraint.AndExpr]()
	serde.RegisterType[constraint.NotExpr]()
	serde.RegisterType[constraint.OrExpr]()
	serde.RegisterType[constraint.SyntaxError]()
	serde.RegisterType[constraint.TagExpr]()
	serde.RegisterType[crypto.Hash]()
	serde.RegisterType[doc.Example]()
	serde.RegisterType[doc.Func]()
	serde.RegisterType[doc.Mode]()
	serde.RegisterType[doc.Note]()
	serde.RegisterType[doc.Package]()
	serde.RegisterType[doc.Type]()
	serde.RegisterType[doc.Value]()
	serde.RegisterType[exec.Cmd]()
	serde.RegisterType[exec.Error]()
	serde.RegisterType[exec.ExitError]()
	serde.RegisterType[float32]()
	serde.RegisterType[float64]()
	serde.RegisterType[fs.FileMode]()
	serde.RegisterType[fs.PathError]()
	serde.RegisterType[int]()
	serde.RegisterType[int16]()
	serde.RegisterType[int32]()
	serde.RegisterType[int64]()
	serde.RegisterType[int8]()
	serde.RegisterType[io.LimitedReader]()
	serde.RegisterType[io.OffsetWriter]()
	serde.RegisterType[io.PipeReader]()
	serde.RegisterType[io.PipeWriter]()
	serde.RegisterType[io.SectionReader]()
	serde.RegisterType[json.Decoder]()
	serde.RegisterType[json.Delim]()
	serde.RegisterType[json.Encoder]()
	serde.RegisterType[json.InvalidUTF8Error]()
	serde.RegisterType[json.InvalidUnmarshalError]()
	serde.RegisterType[json.MarshalerError]()
	serde.RegisterType[json.Number]()
	serde.RegisterType[json.RawMessage]()
	serde.RegisterType[json.SyntaxError]()
	serde.RegisterType[json.UnmarshalFieldError]()
	serde.RegisterType[json.UnmarshalTypeError]()
	serde.RegisterType[json.UnsupportedTypeError]()
	serde.RegisterType[json.UnsupportedValueError]()
	serde.RegisterType[log.Logger]()
	serde.RegisterType[map[*byte][]byte]()
	serde.RegisterType[map[int64]bool]()
	serde.RegisterType[map[int]int]()
	serde.RegisterType[map[string][]byte]()
	serde.RegisterType[map[string]bool]()
	serde.RegisterType[map[string]int]()
	serde.RegisterType[map[string]map[string]bool]()
	serde.RegisterType[map[string]string]()
	serde.RegisterType[map[string]struct{}]()
	serde.RegisterType[map[string]uint64]()
	serde.RegisterType[map[uint64]bool]()
	serde.RegisterType[map[uint64]string]()
	serde.RegisterType[objectpath.Encoder]()
	serde.RegisterType[objectpath.Path]()
	serde.RegisterType[os.File]()
	serde.RegisterType[os.LinkError]()
	serde.RegisterType[os.ProcAttr]()
	serde.RegisterType[os.Process]()
	serde.RegisterType[os.ProcessState]()
	serde.RegisterType[os.SyscallError]()
	serde.RegisterType[packages.Config]()
	serde.RegisterType[packages.Error]()
	serde.RegisterType[packages.ErrorKind]()
	serde.RegisterType[packages.LoadMode]()
	serde.RegisterType[packages.Module]()
	serde.RegisterType[packages.ModuleError]()
	serde.RegisterType[packages.OverlayJSON]()
	serde.RegisterType[packages.Package]()
	serde.RegisterType[parser.Mode]()
	serde.RegisterType[rand.Rand]()
	serde.RegisterType[rand.Zipf]()
	serde.RegisterType[reflect.ChanDir]()
	serde.RegisterType[reflect.Kind]()
	serde.RegisterType[reflect.MapIter]()
	serde.RegisterType[reflect.Method]()
	serde.RegisterType[reflect.SelectCase]()
	serde.RegisterType[reflect.SelectDir]()
	serde.RegisterType[reflect.SliceHeader]()
	serde.RegisterType[reflect.StringHeader]()
	serde.RegisterType[reflect.StructField]()
	serde.RegisterType[reflect.StructTag]()
	serde.RegisterType[reflect.Value]()
	serde.RegisterType[reflect.ValueError]()
	serde.RegisterType[regexp.Regexp]()
	serde.RegisterType[rune]()
	serde.RegisterType[runtime.BlockProfileRecord]()
	serde.RegisterType[runtime.Frame]()
	serde.RegisterType[runtime.Frames]()
	serde.RegisterType[runtime.Func]()
	serde.RegisterType[runtime.MemProfileRecord]()
	serde.RegisterType[runtime.MemStats]()
	serde.RegisterType[runtime.PanicNilError]()
	serde.RegisterType[runtime.Pinner]()
	serde.RegisterType[runtime.StackRecord]()
	serde.RegisterType[runtime.TypeAssertionError]()
	serde.RegisterType[scanner.Error]()
	serde.RegisterType[scanner.ErrorList]()
	serde.RegisterType[scanner.Mode]()
	serde.RegisterType[scanner.Scanner]()
	serde.RegisterType[semver.ByVersion]()
	serde.RegisterType[serde.Deserializer]()
	serde.RegisterType[serde.ID]()
	serde.RegisterType[serde.Serializer]()
	serde.RegisterType[slog.Attr]()
	serde.RegisterType[slog.HandlerOptions]()
	serde.RegisterType[slog.JSONHandler]()
	serde.RegisterType[slog.Kind]()
	serde.RegisterType[slog.Level]()
	serde.RegisterType[slog.LevelVar]()
	serde.RegisterType[slog.Logger]()
	serde.RegisterType[slog.Record]()
	serde.RegisterType[slog.Source]()
	serde.RegisterType[slog.TextHandler]()
	serde.RegisterType[slog.Value]()
	serde.RegisterType[sort.Float64Slice]()
	serde.RegisterType[sort.IntSlice]()
	serde.RegisterType[sort.StringSlice]()
	serde.RegisterType[strconv.NumError]()
	serde.RegisterType[string]()
	serde.RegisterType[strings.Builder]()
	serde.RegisterType[strings.Reader]()
	serde.RegisterType[strings.Replacer]()
	serde.RegisterType[struct{ImportPath string; Export string}]()
	serde.RegisterType[struct{b bool; x any}]()
	serde.RegisterType[struct{base uintptr; end uintptr}]()
	serde.RegisterType[struct{enabled bool; pad [3]byte; needed bool; alignme uint64}]()
	serde.RegisterType[struct{fill uint64; capacity uint64}]()
	serde.RegisterType[struct{ptr interface{}; len int}]()
	serde.RegisterType[struct{tick uint64; i int}]()
	serde.RegisterType[struct{vendor []string; goroot string; gopath []string}]()
	serde.RegisterType[struct{}]()
	serde.RegisterType[sync.Cond]()
	serde.RegisterType[sync.Map]()
	serde.RegisterType[sync.Mutex]()
	serde.RegisterType[sync.Once]()
	serde.RegisterType[sync.Pool]()
	serde.RegisterType[sync.RWMutex]()
	serde.RegisterType[sync.WaitGroup]()
	serde.RegisterType[syntax.EmptyOp]()
	serde.RegisterType[syntax.Error]()
	serde.RegisterType[syntax.ErrorCode]()
	serde.RegisterType[syntax.Flags]()
	serde.RegisterType[syntax.Inst]()
	serde.RegisterType[syntax.InstOp]()
	serde.RegisterType[syntax.Op]()
	serde.RegisterType[syntax.Prog]()
	serde.RegisterType[syntax.Regexp]()
	serde.RegisterType[syscall.Cmsghdr]()
	serde.RegisterType[syscall.Credential]()
	serde.RegisterType[syscall.Dirent]()
	serde.RegisterType[syscall.EpollEvent]()
	serde.RegisterType[syscall.Errno]()
	serde.RegisterType[syscall.FdSet]()
	serde.RegisterType[syscall.Flock_t]()
	serde.RegisterType[syscall.Fsid]()
	serde.RegisterType[syscall.ICMPv6Filter]()
	serde.RegisterType[syscall.IPMreq]()
	serde.RegisterType[syscall.IPMreqn]()
	serde.RegisterType[syscall.IPv6MTUInfo]()
	serde.RegisterType[syscall.IPv6Mreq]()
	serde.RegisterType[syscall.IfAddrmsg]()
	serde.RegisterType[syscall.IfInfomsg]()
	serde.RegisterType[syscall.Inet4Pktinfo]()
	serde.RegisterType[syscall.Inet6Pktinfo]()
	serde.RegisterType[syscall.InotifyEvent]()
	serde.RegisterType[syscall.Iovec]()
	serde.RegisterType[syscall.Linger]()
	serde.RegisterType[syscall.Msghdr]()
	serde.RegisterType[syscall.NetlinkMessage]()
	serde.RegisterType[syscall.NetlinkRouteAttr]()
	serde.RegisterType[syscall.NetlinkRouteRequest]()
	serde.RegisterType[syscall.NlAttr]()
	serde.RegisterType[syscall.NlMsgerr]()
	serde.RegisterType[syscall.NlMsghdr]()
	serde.RegisterType[syscall.ProcAttr]()
	serde.RegisterType[syscall.PtraceRegs]()
	serde.RegisterType[syscall.RawSockaddr]()
	serde.RegisterType[syscall.RawSockaddrAny]()
	serde.RegisterType[syscall.RawSockaddrInet4]()
	serde.RegisterType[syscall.RawSockaddrInet6]()
	serde.RegisterType[syscall.RawSockaddrLinklayer]()
	serde.RegisterType[syscall.RawSockaddrNetlink]()
	serde.RegisterType[syscall.RawSockaddrUnix]()
	serde.RegisterType[syscall.Rlimit]()
	serde.RegisterType[syscall.RtAttr]()
	serde.RegisterType[syscall.RtGenmsg]()
	serde.RegisterType[syscall.RtMsg]()
	serde.RegisterType[syscall.RtNexthop]()
	serde.RegisterType[syscall.Rusage]()
	serde.RegisterType[syscall.Signal]()
	serde.RegisterType[syscall.SockFilter]()
	serde.RegisterType[syscall.SockFprog]()
	serde.RegisterType[syscall.SockaddrInet4]()
	serde.RegisterType[syscall.SockaddrInet6]()
	serde.RegisterType[syscall.SockaddrLinklayer]()
	serde.RegisterType[syscall.SockaddrNetlink]()
	serde.RegisterType[syscall.SockaddrUnix]()
	serde.RegisterType[syscall.SocketControlMessage]()
	serde.RegisterType[syscall.Stat_t]()
	serde.RegisterType[syscall.Statfs_t]()
	serde.RegisterType[syscall.SysProcAttr]()
	serde.RegisterType[syscall.SysProcIDMap]()
	serde.RegisterType[syscall.Sysinfo_t]()
	serde.RegisterType[syscall.TCPInfo]()
	serde.RegisterType[syscall.Termios]()
	serde.RegisterType[syscall.Time_t]()
	serde.RegisterType[syscall.Timespec]()
	serde.RegisterType[syscall.Timeval]()
	serde.RegisterType[syscall.Timex]()
	serde.RegisterType[syscall.Tms]()
	serde.RegisterType[syscall.Ucred]()
	serde.RegisterType[syscall.Ustat_t]()
	serde.RegisterType[syscall.Utimbuf]()
	serde.RegisterType[syscall.Utsname]()
	serde.RegisterType[syscall.WaitStatus]()
	serde.RegisterType[time.Duration]()
	serde.RegisterType[time.Location]()
	serde.RegisterType[time.Month]()
	serde.RegisterType[time.ParseError]()
	serde.RegisterType[time.Ticker]()
	serde.RegisterType[time.Time]()
	serde.RegisterType[time.Timer]()
	serde.RegisterType[time.Weekday]()
	serde.RegisterType[token.File]()
	serde.RegisterType[token.FileSet]()
	serde.RegisterType[token.Pos]()
	serde.RegisterType[token.Position]()
	serde.RegisterType[token.Token]()
	serde.RegisterType[types.ArgumentError]()
	serde.RegisterType[types.Array]()
	serde.RegisterType[types.Basic]()
	serde.RegisterType[types.BasicInfo]()
	serde.RegisterType[types.BasicKind]()
	serde.RegisterType[types.Builtin]()
	serde.RegisterType[types.Chan]()
	serde.RegisterType[types.ChanDir]()
	serde.RegisterType[types.Checker]()
	serde.RegisterType[types.Config]()
	serde.RegisterType[types.Const]()
	serde.RegisterType[types.Context]()
	serde.RegisterType[types.Error]()
	serde.RegisterType[types.Func]()
	serde.RegisterType[types.ImportMode]()
	serde.RegisterType[types.Info]()
	serde.RegisterType[types.Initializer]()
	serde.RegisterType[types.Instance]()
	serde.RegisterType[types.Interface]()
	serde.RegisterType[types.Label]()
	serde.RegisterType[types.Map]()
	serde.RegisterType[types.MethodSet]()
	serde.RegisterType[types.Named]()
	serde.RegisterType[types.Nil]()
	serde.RegisterType[types.Package]()
	serde.RegisterType[types.PkgName]()
	serde.RegisterType[types.Pointer]()
	serde.RegisterType[types.Scope]()
	serde.RegisterType[types.Selection]()
	serde.RegisterType[types.SelectionKind]()
	serde.RegisterType[types.Signature]()
	serde.RegisterType[types.Slice]()
	serde.RegisterType[types.StdSizes]()
	serde.RegisterType[types.Struct]()
	serde.RegisterType[types.Term]()
	serde.RegisterType[types.Tuple]()
	serde.RegisterType[types.TypeAndValue]()
	serde.RegisterType[types.TypeList]()
	serde.RegisterType[types.TypeName]()
	serde.RegisterType[types.TypeParam]()
	serde.RegisterType[types.TypeParamList]()
	serde.RegisterType[types.Union]()
	serde.RegisterType[types.Var]()
	serde.RegisterType[uint]()
	serde.RegisterType[uint16]()
	serde.RegisterType[uint32]()
	serde.RegisterType[uint64]()
	serde.RegisterType[uint8]()
	serde.RegisterType[uintptr]()
	serde.RegisterType[unicode.CaseRange]()
	serde.RegisterType[unicode.Range16]()
	serde.RegisterType[unicode.Range32]()
	serde.RegisterType[unicode.RangeTable]()
	serde.RegisterType[unicode.SpecialCase]()
	serde.RegisterType[unsafe.Pointer]()
}
