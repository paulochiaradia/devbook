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
        Swal.fire({
            title: "Ops....",
            text: "Falha ao criar publicação!",
            icon: "error"
        });
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
        Swal.fire({
            title: "Ops....",
            text: "Falha ao curtir publicação!",
            icon: "error"
        });
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
        Swal.fire({
            title: "Ops....",
            text: "Falha ao descurtir publicação!",
            icon: "error"
        });
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
        Swal.fire({
            title: "Sucesso!",
            text: "Publicacao atualizada com sucesso!",
            icon: "success"
          }).then(function(){
            window.location = "/home"
          });
    }).fail(function(){
        Swal.fire({
            title: "Ops....",
            text: "Falha ao editar publicação!",
            icon: "error"
        });
        
    }).always(function(){
        $('#atualizar-publicacao').prop('disabled', false)
    })
}

function deletarPublicacao(event){
    event.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publcacao?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
      }).then(function(confirmacao){
        if(!confirmacao.value) return
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
            Swal.fire({
                title: "Ops....",
                text: "Falha ao excluir publicação!",
                icon: "error"
            });
        }).always(function(){
            elementoClicado.prop('disabled', false);
        })        
      });

    
}