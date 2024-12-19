package utils

const prompt = `<Contexto> Você é um modelo de criação de questions acadêmicas. Seu foco é criar questions de múltipla escolha com base em um texto que será fornecido. As questions devem avaliar o entendimento do texto, cobrindo seus principais conceitos.
O objetivo é que você crie 10 questions sobre o conteúdo fornecido e retorne as respostas em formato JSON, respeitando a estrutura especificada.

Cada question deve:

Ser clara e diretamente relacionada ao texto.
Ter 4 alternatives de resposta (A, B, C e D).
Ter uma alternativa correta identificada no campo "right_answer".
</Contexto>
<Instruções>
Ao receber o texto, crie 10 questions no seguinte formato JSON:

[
  {
    "question": "Escreva aqui a question completa",
    "alternatives": {
      "a": "Escreva a alternativa A",
      "b": "Escreva a alternativa B",
      "c": "Escreva a alternativa C",
      "d": "Escreva a alternativa D"
    },
    "right_answer": "a"
  },
  {
    "question": "Escreva aqui a segunda question completa",
    "alternatives": {
      "a": "Escreva a alternativa A",
      "b": "Escreva a alternativa B",
      "c": "Escreva a alternativa C",
      "d": "Escreva a alternativa D"
    },
    "right_answer": "b"
  }
]
Exemplo:
Se o texto fornecido for sobre o Ciclo de Krebs, uma question pode ser:

{
  "question": "Qual enzima regula a reação que transforma oxaloacetato em citrato?",
  "alternatives": {
    "a": "Citrato sintase.",
    "b": "Malato desidrogenase.",
    "c": "Fumarase.",
    "d": "Isocitrato desidrogenase."
  },
  "right_answer": "a"
}
<Requisitos adicionais> - Evite questions óbvias; priorize aquelas que realmente testem o conhecimento do texto. - Certifique-se de que o **right_answer** está correto. - Não forneça explicações junto com as respostas, apenas o formato JSON.
</Instruções>
`

const answer = `[
  {
    "question": "Qual é o destino do piruvato em condições aeróbicas após a glicólise?",
    "alternatives": {
      "a": "Transformação em lactato.",
      "b": "Transformação em etanol.",
      "c": "Descarboxilação oxidativa para formar acetil-CoA.",
      "d": "Conversão direta em oxaloacetato."
    },
    "right_answer": "c"
  },
  {
    "question": "Qual é a função principal do ciclo de Krebs?",
    "alternatives": {
      "a": "Produzir glicose.",
      "b": "Oxidar o acetil-CoA e produzir coenzimas reduzidas.",
      "c": "Armazenar energia na forma de glicogênio.",
      "d": "Converter piruvato em lactato."
    },
    "right_answer": "b"
  },
  {
    "question": "Que tipo de reação ocorre na mitocôndria envolvendo o piruvato antes de entrar no ciclo de Krebs?",
    "alternatives": {
      "a": "Carboxilação.",
      "b": "Oxidação.",
      "c": "Descarboxilação oxidativa.",
      "d": "Redução."
    },
    "right_answer": "c"
  },
  {
    "question": "Qual é o papel das reações anapleróticas no ciclo de Krebs?",
    "alternatives": {
      "a": "Gerar oxigênio para o ciclo.",
      "b": "Manter a concentração dos intermediários do ciclo.",
      "c": "Consumir produtos finais do ciclo.",
      "d": "Promover a redução de NAD+."
    },
    "right_answer": "b"
  },
  {
    "question": "O que acontece com NADH e FADH2 produzidos no ciclo de Krebs?",
    "alternatives": {
      "a": "São utilizados diretamente na síntese de ATP.",
      "b": "São oxidados na cadeia respiratória para gerar ATP.",
      "c": "São convertidos em CO2.",
      "d": "São armazenados na mitocôndria."
    },
    "right_answer": "b"
  },
  {
    "question": "Qual das seguintes enzimas está envolvida na transformação do piruvato em acetil-CoA?",
    "alternatives": {
      "a": "Citrato sintase.",
      "b": "Isocitrato desidrogenase.",
      "c": "Piruvato desidrogenase.",
      "d": "Fumarase."
    },
    "right_answer": "c"
  },
  {
    "question": "Quais são os principais controladores alostéricos negativos da piruvato desidrogenase?",
    "alternatives": {
      "a": "ATP e succinato.",
      "b": "NADH e acetil-CoA.",
      "c": "CO2 e ADP.",
      "d": "FADH2 e glicose."
    },
    "right_answer": "b"
  },
  {
    "question": "Que função desempenham as enzimas citrato sintase, isocitrato desidrogenase e α-cetoglutarato desidrogenase no ciclo de Krebs?",
    "alternatives": {
      "a": "Elas convertem CO2 em NADH.",
      "b": "Regulam o fluxo metabólico do ciclo.",
      "c": "São responsáveis pela síntese de ATP.",
      "d": "Armazenam energia na forma de ADP."
    },
    "right_answer": "b"
  },
  {
    "question": "Por que o ciclo de Krebs é classificado como um ciclo anfibólico?",
    "alternatives": {
      "a": "Somente porque gera energia.",
      "b": "Porque participa de vias catabólicas e anabólicas.",
      "c": "Porque ocorre em ambiente aquático.",
      "d": "Porque não envolve nenhuma oxidação."
    },
    "right_answer": "b"
  },
  {
    "question": "Qual é o papel do piruvato carboxilase no ciclo de Krebs?",
    "alternatives": {
      "a": "Descarboxilar piruvato em acetil-CoA.",
      "b": "Reduzir NAD+ a NADH.",
      "c": "Carboxilar piruvato para formar oxaloacetato.",
      "d": "Oxidar CO2 para formar piruvato."
    },
    "right_answer": "c"
  }
]`

func GetPrompt() string {
	return prompt
}
func GetAnswer() string {
	return answer
}
