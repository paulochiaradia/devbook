$('#nova-publicacao').on('submit', criarPublicacao);
$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao) ;
$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);


function criarPublicacao(event) {
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
    
    elementoClicado.prop('disabled', true);
    $.ajax({
        url :`/publicacoes/${publicacaoID}/curtir`,
        method: "POST"
    }).done(function(){
        const contadorDeCurtidas = elementoClicado.next("span");
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas+1);

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');
    }).fail(function(){
        alert("Erro ao curtir a publicacao");
    }).always(function(){
        elementoClicado.prop('disabled', false);
    })
}

function descurtirPublicacao(event) {
    event.preventDefault();

    const elementoClicado = $(event.target);
    const publicacaoDiv = elementoClicado.closest(".jumbotron-style");
    const publicacaoID = publicacaoDiv.data("publicacao-id");
    
    elementoClicado.prop('disabled', true);
    $.ajax({
        url :`/publicacoes/${publicacaoID}/descurtir`,
        method: "POST"
    }).done(function(){
        const contadorDeCurtidas = elementoClicado.next("span");
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas-1);

        elementoClicado.removeClass('descurtir-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');
    }).fail(function(){
        alert("Erro ao descurtir a publicacao")
    }).always(function(){
        elementoClicado.prop('disabled', false);
    })
}

function atualizarPublicacao(){
    $(this).prop('disabled', true)
    const publicacaoID = $(this).data('publicacao-id');
    $.ajax({
        url: `/publicacoes/${publicacaoID}`,
        method : "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function(){
        alert("Publicacao editada com sucesso")
    }).fail(function(){
        alert("Erro ao editar publicacao")
    }).always(function(){
        $('#atualizar-publicacao').prop('disabled', false)
    })
}

function deletarPublicacao(event){
    event.preventDefault();

    const elementoClicado = $(event.target);
    const publicacao = elementoClicado.closest(".jumbotron-style");
    const publicacaoID = publicacao.data("publicacao-id");
    
    elementoClicado.prop('disabled', true);
    $.ajax({
        url :`/publicacoes/${publicacaoID}`,
        method: "DELETE"
    }).done(function(){
        publicacao.fadeOut("slow", function(){
            $(this).remove();
        });
    }).fail(function(){
        alert("Erro ao excluir a publicacao")
    }).always(function(){
        elementoClicado.prop('disabled', false);
    })
}