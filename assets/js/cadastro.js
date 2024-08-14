$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("aí papai");

    if($('#senha').val() != $('#confirmar-senha').val()){
        alert("tá errado ô");
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
    })
}