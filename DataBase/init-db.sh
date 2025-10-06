#!/bin/bash
set -e

echo "🏗️ Iniciando creación de tablas..."

INIT_FILE="/database/init-order.txt"

if [ ! -f "$INIT_FILE" ]; then
  echo "❌ No se encontró el archivo init-order.txt"
  exit 1
fi

# Ejecutar tablas
while IFS= read -r line; do
  [[ -z "$line" || "$line" =~ ^# ]] && continue
  FILE_PATH="/database/$line"
  if [ -f "$FILE_PATH" ]; then
    echo "🧱 Ejecutando: $FILE_PATH"
    psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f "$FILE_PATH"
  else
    echo "⚠️ Archivo no encontrado: $FILE_PATH"
  fi
done < "$INIT_FILE"

# Insertar datos
echo "🌱 Insertando datos iniciales..."
while IFS= read -r line; do
  [[ -z "$line" || "$line" =~ ^# ]] && continue
  DATA_FILE=$(echo "$line" | sed 's|Tables/|Data/|g')
  DATA_PATH="/database/$DATA_FILE"
  if [ -f "$DATA_PATH" ]; then
    echo "📥 Insertando: $DATA_PATH"
    psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f "$DATA_PATH"
  fi
done < "$INIT_FILE"

echo "🎉 Base de datos inicializada correctamente."
