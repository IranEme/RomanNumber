# Roman Numerals Converter

Este es un proyecto en Go que permite convertir números romanos a enteros y viceversa. La aplicación toma una entrada del usuario, ya sea un número romano o un número entero, y proporciona la conversión correspondiente.

## Funcionalidades

- **Conversión de números romanos a enteros**: Convierte una cadena de texto con un número romano válido en su equivalente en número entero.
- **Conversión de enteros a números romanos**: Convierte un número entero en su representación en números romanos.

## Estructura del Código

El programa se compone de las siguientes funciones principales:

- **`isValidRoman(s string) bool`**: Verifica si una cadena es un número romano válido utilizando expresiones regulares.
- **`RomanToInt(s string) (int, error)`**: Convierte un número romano en su valor entero correspondiente.
- **`IntToRoman(num int) string`**: Convierte un número entero en su representación en números romanos.
- **`main()`**: Solicita una entrada del usuario, detecta si es un número entero o romano, y llama a la función correspondiente para hacer la conversión.

## Requisitos

- **Go 1.x** o superior.

## Instalación

1. Clona el repositorio en tu máquina local:

   ```bash
   git clone https://github.com/tu-usuario/roman-numerals-converter.git
   ```

2. Navega al directorio del proyecto:

   ```bash
   cd roman-numerals-converter
   ```

3. Ejecuta el proyecto:

   ```bash
   go run main.go
   ```

## Uso

Al ejecutar el programa, se te pedirá ingresar un número o un número romano:

- Si introduces un número entero, el programa mostrará su equivalente en número romano.
- Si introduces un número romano, el programa mostrará su equivalente en número entero.

### Ejemplos

**Entrada**: `123`

**Salida**: `CXXIII`

**Entrada**: `XIV`

**Salida**: `14`

## Contribución

Si deseas contribuir al proyecto, por favor, realiza un fork del repositorio, haz tus cambios y envía un pull request.
