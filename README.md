# 🚀 Taller Práctico: Prompting para Desarrollo y Project Management con Gemini CLI

¡Bienvenidos al taller! En esta sesión práctica aprenderemos a utilizar **Gemini CLI** como un asistente impulsado por IA directamente desde nuestra terminal. Exploraremos cómo la inteligencia artificial generativa puede acelerar flujos de trabajo de modernización de aplicaciones, DevOps, refactorización de código y gestión de proyectos.

## 📋 Prerrequisitos

Antes de comenzar el taller, asegúrate de contar con lo siguiente en tu entorno local:

1. **Git:** Para clonar este repositorio.
2. **Gemini CLI:** Instalado y autenticado con tus credenciales. 
3. **Go y Java (Opcional pero recomendado):** Para ejecutar y probar el código de los retos 1, 3 y 4 localmente.
4. **Un editor de código o terminal:** VS Code, IntelliJ, o simplemente tu terminal favorita.

## ⏱️ Dinámica del Taller

El taller está dividido en 4 retos prácticos. Cada reto está diseñado para resolverse en aproximadamente **15 minutos**.

> **⚠️ Regla de oro:** Para que Gemini CLI tenga el contexto correcto sin "contaminarse" de otros ejercicios, asegúrate de navegar a la carpeta específica de cada reto antes de ejecutar tus prompts.

---

## 📂 Estructura de los Retos

### ☕ [Reto 1: Java] Comprensión de Código y Contenerización
*   **Directorio:** `cd reto-1-java`
*   **Misión:** Analizar un archivo de Java legado (`LegacyOrderProcessor.java`), generar su documentación (`README.md`) y crear un `Dockerfile` optimizado para Cloud Run utilizando Gemini.

### 📊 [Reto 2: PM] Planificación de Migración y Métricas
*   **Directorio:** `cd reto-2-pm`
*   **Misión:** Utilizar el documento `contexto_arquitectura.md` para que Gemini actúe como Technical Project Manager, diseñando un plan de migración a microservicios en GKE y definiendo métricas DORA para medir el éxito.

### 🐹 [Reto 3: Go] Debugging y Refactorización
*   **Directorio:** `cd reto-3-go`
*   **Misión:** Detectar vulnerabilidades de concurrencia y cuellos de botella en un microservicio en Golang (`main.go`), refactorizar el código con buenas prácticas y generar sus pruebas unitarias.

### ⚙️ [Reto 4: CI/CD] Definición de Pipelines y Backlog Ágil
*   **Directorio:** `cd reto-4-cicd`
*   **Misión:** Integrar el lado técnico y el ágil: generar un archivo `cloudbuild.yaml` para automatizar pruebas y construcción de contenedores, y luego redactar la historia de usuario correspondiente para el próximo sprint.

