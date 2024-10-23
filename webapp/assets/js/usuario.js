$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);
$('#editar-usuario').on('submit', editar);
$('#atualizar-senha').on('submit', atualizarSenha);

function pararDeSeguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuario/${usuarioId}/parar-de-seguir`,
        method: "POST"
    }).done(function() {
        window.location = `/usuario/${usuarioId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário!", "error");
        $('#parar-de-seguir').prop('disabled', false);
    });
}

function seguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuario/${usuarioId}/seguir`,
        method: "POST"
    }).done(function() {
        window.location = `/usuario/${usuarioId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao seguir o usuário!", "error");
        $('#seguir').prop('disabled', false);
    });
}

function editar(event){
    event.preventDefault();
    $.ajax({
        url: `/editar-usuario`,
        method: "PUT",
        data: {
            nome: $(`#nome`).val(),
            email: $(`#email`).val(),
            nick: $(`#nick`).val(),
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Usuario atualizado com sucesso!", "success")
        .then(function(){
            window.location = `/perfil`;
        });
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error");
    });
}

function atualizarSenha(event){
    event.preventDefault();

    if($(`#nova-senha`).val() !=$(`#confirmar-senha`).val()){
        Swal.Fire("Ops...", "As senhas nao sao iguais!", "warning");
        return;
    }

    $.ajax({
        url: `/atualizar-senha`,
        method: "POST",
        data: {
            atual: $(`#senha-atual`).val(),
            nova: $(`#nova-senha`).val(),
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Senha atualizada com sucesso!", "success")
        .then(function(){
            window.location = `/perfil`;
        });
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar a senha!", "error");
    });

}