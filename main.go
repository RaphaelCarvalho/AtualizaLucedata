package main

import (
	"fmt"
	"io"
	"os"
)

func copiarArquivo(origem, destino string) error {
	// Abre o arquivo de origem para leitura
	arquivoOrigem, err := os.Open(origem)
	if err != nil {
		return err
	}
	defer arquivoOrigem.Close()

	// Cria o arquivo de destino
	arquivoDestino, err := os.Create(destino)
	if err != nil {
		return err
	}
	defer arquivoDestino.Close()

	// Copia o conteúdo do arquivo de origem para o arquivo de destino
	_, err = io.Copy(arquivoDestino, arquivoOrigem)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Especifica os caminhos dos arquivos de origem e destino
	origens := []string{
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Administrativo\\Administrativo.exe",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Administrativo\\versao.txt",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Almoxarifado\\Almoxarifado.exe",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Almoxarifado\\versao.txt",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Compras\\Compras.exe",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Compras\\versao.txt",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Farmacia\\Farmacia.exe",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Farmacia\\versao.txt",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Faturamento\\Faturamento.exe",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Faturamento\\versao.txt",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Recepcao\\Recepcao.exe",
		"\\\\192.168.15.2\\Lucedata\\atualizacao\\Recepcao\\versao.txt",
	}

	destinos := []string{
		"C:\\SGH\\Administrativo\\Administrativo.exe",
		"C:\\SGH\\Administrativo\\versao.txt",
		"C:\\SGH\\Almoxarifado\\Almoxarifado.exe",
		"C:\\SGH\\Almoxarifado\\versao.txt",
		"C:\\SGH\\Compras\\Compras.exe",
		"C:\\SGH\\Compras\\versao.txt",
		"C:\\SGH\\Farmacia\\Farmacia.exe",
		"C:\\SGH\\Farmacia\\versao.txt",
		"C:\\SGH\\Faturamento\\Faturamento.exe",
		"C:\\SGH\\Faturamento\\versao.txt",
		"C:\\SGH\\Recepcao\\Recepcao.exe",
		"C:\\SGH\\Recepcao\\versao.txt",
	}

	// Realiza as operações de cópia
	for i, origem := range origens {
		err := copiarArquivo(origem, destinos[i])
		if err != nil {
			fmt.Printf("Erro durante as operações de cópia: %v\n", err)
			return
		}
		fmt.Printf("Arquivo copiado com sucesso: %v\n", origem)
	}

	fmt.Println("Operações de cópia concluídas com sucesso.")
}
