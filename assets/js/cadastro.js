$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();

    if($('#senha').val() != $('#confirmar-senha').val()){
        alert("Confirme a senha corretamente");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome:  $('#nome').val(),
            email: $('#email').val(),
            nick:  $('#nick').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {
        alert("Usuário cadastrado com sucesso!")
    }).fail(function() {
        alert("Erro ao cadastrar o usuário!")
    });
}
