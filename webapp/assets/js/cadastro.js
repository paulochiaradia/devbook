$('#formulario-cadastro').on('submit',criarUsuario);

function criarUsuario(event){
    event.preventDefault();
    console.log("Esta funcionando");

    if ($('#senha').val() !== $('#confirmar-senha').val()){
        alert('As senhas não conferem');
    }
    $.ajax({
        url: '/usuarios',
        method: 'POST',
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        },
    }).done(function (){
        alert("Usuario Cadastrado com Sucesso")
    }).fail(function (){
        console.log("erro")
        alert("Erro ao cadastrar usuario")
    })
}