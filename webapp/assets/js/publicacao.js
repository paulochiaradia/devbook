$('#nova-publicacao').on('submit', criariPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao)

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
        window.location = "/home";
    }).fail(function(){
        alert('Falha ao criar publicação!');
    });

}

function curtirPublicacao(event) {
    event.preventDefault();

    const elementoClicado = $(event.target);
    const publicacaoDiv = elementoClicado.closest(".jumbotron-style");
    const publicacaoID = publicacaoDiv.data("publicacao-id");
    
    $.ajax({
        url :`/publicacoes/${publicacaoID}/curtir`,
        method: "POST"
    }).done(function(){
        alert("Publicaco curtida");
    }).fail(function(){
        alert("Erro ao curtir a publicacao")
    })


}