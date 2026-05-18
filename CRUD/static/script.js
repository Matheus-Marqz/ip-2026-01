const form = document.getElementById("formPaciente")
const pacientesDiv = document.getElementById("pacientes")
const mensagem = document.getElementById("mensagem")

form.addEventListener("submit", async function (event) {
	event.preventDefault()

	const id = document.getElementById("id").value
	const paciente = pegarDadosFormulario()

	if (id) {
		await fetch("/pacientes/update?id=" + id, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(paciente),
		})

		mostrarMensagem("Paciente atualizado")
	} else {
		await fetch("/pacientes/create", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(paciente),
		})

		mostrarMensagem("Paciente cadastrado")
	}

	limparFormulario()
	listarPacientes()
})

function pegarDadosFormulario() {
	return {
		name: document.getElementById("name").value,
		fone: document.getElementById("fone").value,
		data_nasc: document.getElementById("data_nasc").value,
		sexo: document.getElementById("sexo").value,
		num_convenio: document.getElementById("num_convenio").value,
	}
}

function limparFormulario() {
	form.reset()
	document.getElementById("id").value = ""
}

async function listarPacientes() {
	const resposta = await fetch("/pacientes/read")
	const pacientes = await resposta.json()

	pacientesDiv.innerHTML = ""

	if (pacientes.length === 0) {
		pacientesDiv.innerHTML = "<p>Nenhum paciente cadastrado.</p>"
		return
	}

	pacientes.forEach(function (paciente) {
		const item = document.createElement("div")
		item.className = "paciente"

		item.innerHTML = `
			<div>
				<strong>${paciente.name}</strong>
				<span>ID: ${paciente.id}</span>
				<span>Telefone: ${paciente.fone}</span>
				<span>Nascimento: ${paciente.data_nasc}</span>
				<span>Sexo: ${paciente.sexo}</span>
				<span>Convenio: ${paciente.num_convenio}</span>
			</div>

			<div class="botoes">
				<button onclick='editarPaciente(${JSON.stringify(paciente)})'>Editar</button>
				<button class="botao-secundario" onclick="deletarPaciente(${paciente.id})">Deletar</button>
			</div>
		`

		pacientesDiv.appendChild(item)
	})
}

function editarPaciente(paciente) {
	document.getElementById("id").value = paciente.id
	document.getElementById("name").value = paciente.name
	document.getElementById("fone").value = paciente.fone
	document.getElementById("data_nasc").value = paciente.data_nasc.substring(0, 10)
	document.getElementById("sexo").value = paciente.sexo
	document.getElementById("num_convenio").value = paciente.num_convenio
}

async function deletarPaciente(id) {
	await fetch("/pacientes/delete?id=" + id, {
		method: "DELETE",
	})

	mostrarMensagem("Paciente deletado")
	listarPacientes()
}

function mostrarMensagem(texto) {
	mensagem.textContent = texto
	setTimeout(function () {
		mensagem.textContent = ""
	}, 2500)
}

listarPacientes()
