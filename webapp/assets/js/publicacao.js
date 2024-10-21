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
    
    elementoClicado.prop('disabled', true)
    $.ajax({
        url :`/publicacoes/${publicacaoID}/curtir`,
        method: "POST"
    }).done(function(){
        const contadorDeCurtidas = elementoClicado.next("span");
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas+1)
    }).fail(function(){
        alert("Erro ao curtir a publicacao")
    }).always(function(){
        elementoClicado.prop('disabled', false)
    })


}