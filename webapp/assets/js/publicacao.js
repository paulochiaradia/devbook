$('#nova-publicacao').on('submit', criariPublicacao);

function criariPublicacao(event) {
    event.preventDefault();
  
    $.ajax({
       url :"/publicacoes",
       method: "POST",
       data :{
        titulo: $('#titulo').val(),
        conteudo: $('#conteudo').val()
       } 
    }).done(function(){
        alert('Publicação criada com sucesso!');
    }).fail(function(){
        alert('Falha ao criar publicação!');
    });

}
