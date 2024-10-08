package entity

type Universal struct {
	Projetos_UNV Projetos_UNV
	Significado  Significado
}

type Projetos_UNV struct {
	ConvencaoPastas ConvencaoPastas
	Entidade        Entidade
}

type Entidade struct {
	OrdemFluxoEntidade []OrdemFluxoEntidade
	CasosDeUsoEntidade CasosDeUsoEntidade
}

type OrdemFluxoEntidade struct {
	Contrato_Formato_Entidade OrdemFluxoEntidadeProps
	Usecase_AcaoCasoDeUso     OrdemFluxoEntidadeProps
	Seed                      OrdemFluxoEntidadeProps
	Tester                    OrdemFluxoEntidadeProps
	Test                      OrdemFluxoEntidadeProps
	Index_Modulo              OrdemFluxoEntidadeProps
	Main_Root                 OrdemFluxoEntidadeProps
}
type OrdemFluxoEntidadeProps struct {
	Ordem      int
	Finalidade Significado
}

type CasosDeUsoEntidade struct {
	Crud []Crud
}

type Crud struct {
	Criar   string
	Listar  string
	Deletar string
}

type Nomeacoes_Pastas = int

const (
	minuscula Nomeacoes_Pastas = iota
	maiusculo
	camelCase
	snackCase
	snackCase_Com_Anderline
)

type Nomeacoes_Arquivos = int

const (
	entidade_acao Nomeacoes_Arquivos = iota
)

type ConvencaoPastas struct {
	NomePastaEntidade Nomeacoes_Pastas
	Nome_Arquivo      Nomeacoes_Arquivos
}

type Significado struct {
	Nome           string
	Significado    string
	Aplicabilidade []string
}
