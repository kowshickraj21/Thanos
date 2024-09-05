package lexer

import "fmt"

type TokenType int

const (
	EOF TokenType = iota
	Identifier
	Keyword
	Number
	CharConst
	StringLiteral

	Plus           // +
	Minus          // -
	Multiply       // *
	Divide         // /
	Modulus        // %
	Increment      // ++
	Decrement      // --
	Assignment     // =
	PlusAssign     // +=
	MinusAssign    // -=
	MultiplyAssign // *=
	DivideAssign   // /=
	ModulusAssign  // %=

	Equal        // ==
	NotEqual     // !=
	GreaterThan  // >
	LessThan     // <
	GreaterEqual // >=
	LessEqual    // <=

	LogicalAnd // &&
	LogicalOr  // ||
	LogicalNot // !
	BitwiseAnd // &
	BitwiseOr  // |
	BitwiseXor // ^
	BitwiseNot // ~

	LeftShift        // <<
	RightShift       // >>
	AndAssign        // &=
	OrAssign         // |=
	XorAssign        // ^=
	LeftShiftAssign  // <<=
	RightShiftAssign // >>=

	ParenOpen    // (
	ParenClose   // )
	BraceOpen    // {
	BraceClose   // }
	BracketOpen  // [
	BracketClose // ]

	Semicolon    // ;
	Colon        // :
	Comma        // ,
	Dot          // .
	QuestionMark // ?
	Hash         // #
	DoubleHash   // ##
	Arrow        // ->
	Ellipsis     // ...

	Preprocessor // Preprocessor tokens like #define, #include
	TypeDef      // typedef
	Extern       // extern
	Static       // static
	Const        // const
	Volatile     // volatile
	Struct       // struct
	Union        // union
	Enum         // enum

	Signed   // signed
	Unsigned // unsigned

	Sizeof   // sizeof
	Goto     // goto
	Return   // return
	Break    // break
	Continue // continue

	Switch  // switch
	Case    // case
	Default // default
	Do      // do
	While   // while
	For     // for
	If      // if
	Else    // else

	LineComment     // //
	BlockComment    // /* */
	Backslash       // \\
	SingleQuote     // \'
	DoubleQuote     // \"
	QuestionMarkLit // \?

	OctalNumber // \0 to \377
	HexNumber   // \x0 to \xFF
)

type Token struct {
	Type  TokenType
	Value string
}

func (token Token) isAny(expectedTokens ...TokenType) bool {
	for _,expected := range expectedTokens {
		if expected == token.Type{
			return true;
		}
	}
	return false;
}

func (token Token) Debug() {
	if token.isAny(Identifier,StringLiteral,Number,CharConst) {
		fmt.Printf("%s ,(%s)\n", tokenToString(token.Type), token.Value)
	}else{
		fmt.Printf("%s ,()\n", tokenToString(token.Type))
	}
}

func newToken(Type TokenType,value string) Token{
	return Token{
		Type, value,
	}
}

func tokenToString(t TokenType) string {
	switch t {
	case EOF:
		return "EOF"
	case Identifier:
		return "Identifier"
	case Keyword:
		return "Keyword"
	case Number:
		return "Number"
	case CharConst:
		return "CharConst"
	case StringLiteral:
		return "StringLiteral"
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Multiply:
		return "Multiply"
	case Divide:
		return "Divide"
	case Modulus:
		return "Modulus"
	case Increment:
		return "Increment"
	case Decrement:
		return "Decrement"
	case Assignment:
		return "Assignment"
	case PlusAssign:
		return "PlusAssign"
	case MinusAssign:
		return "MinusAssign"
	case MultiplyAssign:
		return "MultiplyAssign"
	case DivideAssign:
		return "DivideAssign"
	case ModulusAssign:
		return "ModulusAssign"
	case Equal:
		return "Equal"
	case NotEqual:
		return "NotEqual"
	case GreaterThan:
		return "GreaterThan"
	case LessThan:
		return "LessThan"
	case GreaterEqual:
		return "GreaterEqual"
	case LessEqual:
		return "LessEqual"
	case LogicalAnd:
		return "LogicalAnd"
	case LogicalOr:
		return "LogicalOr"
	case LogicalNot:
		return "LogicalNot"
	case BitwiseAnd:
		return "BitwiseAnd"
	case BitwiseOr:
		return "BitwiseOr"
	case BitwiseXor:
		return "BitwiseXor"
	case BitwiseNot:
		return "BitwiseNot"
	case LeftShift:
		return "LeftShift"
	case RightShift:
		return "RightShift"
	case AndAssign:
		return "AndAssign"
	case OrAssign:
		return "OrAssign"
	case XorAssign:
		return "XorAssign"
	case LeftShiftAssign:
		return "LeftShiftAssign"
	case RightShiftAssign:
		return "RightShiftAssign"
	case ParenOpen:
		return "ParenOpen"
	case ParenClose:
		return "ParenClose"
	case BraceOpen:
		return "BraceOpen"
	case BraceClose:
		return "BraceClose"
	case BracketOpen:
		return "BracketOpen"
	case BracketClose:
		return "BracketClose"
	case Semicolon:
		return "Semicolon"
	case Colon:
		return "Colon"
	case Comma:
		return "Comma"
	case Dot:
		return "Dot"
	case QuestionMark:
		return "QuestionMark"
	case Hash:
		return "Hash"
	case DoubleHash:
		return "DoubleHash"
	case Arrow:
		return "Arrow"
	case Ellipsis:
		return "Ellipsis"
	case Preprocessor:
		return "Preprocessor"
	case TypeDef:
		return "TypeDef"
	case Extern:
		return "Extern"
	case Static:
		return "Static"
	case Const:
		return "Const"
	case Volatile:
		return "Volatile"
	case Struct:
		return "Struct"
	case Union:
		return "Union"
	case Enum:
		return "Enum"
	case Signed:
		return "Signed"
	case Unsigned:
		return "Unsigned"
	case Sizeof:
		return "Sizeof"
	case Goto:
		return "Goto"
	case Return:
		return "Return"
	case Break:
		return "Break"
	case Continue:
		return "Continue"
	case Switch:
		return "Switch"
	case Case:
		return "Case"
	case Default:
		return "Default"
	case Do:
		return "Do"
	case While:
		return "While"
	case For:
		return "For"
	case If:
		return "If"
	case Else:
		return "Else"
	case LineComment:
		return "LineComment"
	case BlockComment:
		return "BlockComment"
	case Backslash:
		return "Backslash"
	case SingleQuote:
		return "SingleQuote"
	case DoubleQuote:
		return "DoubleQuote"
	case QuestionMarkLit:
		return "QuestionMarkLit"
	case OctalNumber:
		return "OctalNumber"
	case HexNumber:
		return "HexNumber"
	default:
		return "Unknown"
	}
}
