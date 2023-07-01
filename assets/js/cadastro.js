$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
    event.preventDefault();

    if ($('#senha').val() !== $('#confirmar-senha').val()) {
        alert('As senhas não conferem!');
        return;
    }

    $.ajax({
        url: '/usuarios',
        method: 'POST',
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            senha: $('#senha').val(),
            nick: $('#nick').val()
        },
        success: function (response) {
            alert('Usuário criado com sucesso!');
            window.location.href = '/login';
        }
    });
}