package utils

const prompt = `<Contexto> Você é um modelo de criação de perguntas acadêmicas. Seu foco é criar perguntas de múltipla escolha com base em um texto que será fornecido. As perguntas devem avaliar o entendimento do texto, cobrindo seus principais conceitos.
O objetivo é que você crie 10 perguntas sobre o conteúdo fornecido e retorne as respostas em formato JSON, respeitando a estrutura especificada.

Cada pergunta deve:

Ser clara e diretamente relacionada ao texto.
Ter 4 alternativas de resposta (A, B, C e D).
Ter uma alternativa correta identificada no campo "gabarito".
</Contexto>
<Instruções>
Ao receber o texto, crie 10 perguntas no seguinte formato JSON:

[
  {
    "pergunta": "Escreva aqui a pergunta completa",
    "alternativas": {
      "a": "Escreva a alternativa A",
      "b": "Escreva a alternativa B",
      "c": "Escreva a alternativa C",
      "d": "Escreva a alternativa D"
    },
    "gabarito": "a"
  },
  {
    "pergunta": "Escreva aqui a segunda pergunta completa",
    "alternativas": {
      "a": "Escreva a alternativa A",
      "b": "Escreva a alternativa B",
      "c": "Escreva a alternativa C",
      "d": "Escreva a alternativa D"
    },
    "gabarito": "b"
  }
]
Exemplo:
Se o texto fornecido for sobre o Ciclo de Krebs, uma pergunta pode ser:

{
  "pergunta": "Qual enzima regula a reação que transforma oxaloacetato em citrato?",
  "alternativas": {
    "a": "Citrato sintase.",
    "b": "Malato desidrogenase.",
    "c": "Fumarase.",
    "d": "Isocitrato desidrogenase."
  },
  "gabarito": "a"
}
<Requisitos adicionais> - Evite perguntas óbvias; priorize aquelas que realmente testem o conhecimento do texto. - Certifique-se de que o **gabarito** está correto. - Não forneça explicações junto com as respostas, apenas o formato JSON.
</Instruções>
`

const answer = `[
  {
    "pergunta": "Qual é o destino do piruvato em condições aeróbicas após a glicólise?",
    "alternativas": {
      "a": "Transformação em lactato.",
      "b": "Transformação em etanol.",
      "c": "Descarboxilação oxidativa para formar acetil-CoA.",
      "d": "Conversão direta em oxaloacetato."
    },
    "gabarito": "c"
  },
  {
    "pergunta": "Qual é a função principal do ciclo de Krebs?",
    "alternativas": {
      "a": "Produzir glicose.",
      "b": "Oxidar o acetil-CoA e produzir coenzimas reduzidas.",
      "c": "Armazenar energia na forma de glicogênio.",
      "d": "Converter piruvato em lactato."
    },
    "gabarito": "b"
  },
  {
    "pergunta": "Que tipo de reação ocorre na mitocôndria envolvendo o piruvato antes de entrar no ciclo de Krebs?",
    "alternativas": {
      "a": "Carboxilação.",
      "b": "Oxidação.",
      "c": "Descarboxilação oxidativa.",
      "d": "Redução."
    },
    "gabarito": "c"
  },
  {
    "pergunta": "Qual é o papel das reações anapleróticas no ciclo de Krebs?",
    "alternativas": {
      "a": "Gerar oxigênio para o ciclo.",
      "b": "Manter a concentração dos intermediários do ciclo.",
      "c": "Consumir produtos finais do ciclo.",
      "d": "Promover a redução de NAD+."
    },
    "gabarito": "b"
  },
  {
    "pergunta": "O que acontece com NADH e FADH2 produzidos no ciclo de Krebs?",
    "alternativas": {
      "a": "São utilizados diretamente na síntese de ATP.",
      "b": "São oxidados na cadeia respiratória para gerar ATP.",
      "c": "São convertidos em CO2.",
      "d": "São armazenados na mitocôndria."
    },
    "gabarito": "b"
  },
  {
    "pergunta": "Qual das seguintes enzimas está envolvida na transformação do piruvato em acetil-CoA?",
    "alternativas": {
      "a": "Citrato sintase.",
      "b": "Isocitrato desidrogenase.",
      "c": "Piruvato desidrogenase.",
      "d": "Fumarase."
    },
    "gabarito": "c"
  },
  {
    "pergunta": "Quais são os principais controladores alostéricos negativos da piruvato desidrogenase?",
    "alternativas": {
      "a": "ATP e succinato.",
      "b": "NADH e acetil-CoA.",
      "c": "CO2 e ADP.",
      "d": "FADH2 e glicose."
    },
    "gabarito": "b"
  },
  {
    "pergunta": "Que função desempenham as enzimas citrato sintase, isocitrato desidrogenase e α-cetoglutarato desidrogenase no ciclo de Krebs?",
    "alternativas": {
      "a": "Elas convertem CO2 em NADH.",
      "b": "Regulam o fluxo metabólico do ciclo.",
      "c": "São responsáveis pela síntese de ATP.",
      "d": "Armazenam energia na forma de ADP."
    },
    "gabarito": "b"
  },
  {
    "pergunta": "Por que o ciclo de Krebs é classificado como um ciclo anfibólico?",
    "alternativas": {
      "a": "Somente porque gera energia.",
      "b": "Porque participa de vias catabólicas e anabólicas.",
      "c": "Porque ocorre em ambiente aquático.",
      "d": "Porque não envolve nenhuma oxidação."
    },
    "gabarito": "b"
  },
  {
    "pergunta": "Qual é o papel do piruvato carboxilase no ciclo de Krebs?",
    "alternativas": {
      "a": "Descarboxilar piruvato em acetil-CoA.",
      "b": "Reduzir NAD+ a NADH.",
      "c": "Carboxilar piruvato para formar oxaloacetato.",
      "d": "Oxidar CO2 para formar piruvato."
    },
    "gabarito": "c"
  }
]`

func GetPrompt() string {
	return prompt
}
func GetAnswer() string {
	return answer
}
