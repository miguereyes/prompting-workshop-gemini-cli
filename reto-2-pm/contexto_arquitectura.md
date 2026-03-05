# Estado actual del Sistema Core
- **Arquitectura:** Monolítica.
- **Lenguaje/Framework:** Java EE corriendo sobre servidores de aplicaciones (WebLogic).
- **Base de Datos:** Oracle DB on-premise, altamente acoplada.
- **Despliegue:** Manual, toma 4 horas en ventana de mantenimiento nocturno.
- **Problemas principales:** 
  - Tiempos de inactividad durante picos de tráfico.
  - El equipo de desarrollo tarda 3 semanas en liberar una nueva funcionalidad.
- **Objetivo del negocio:** Migrar a una arquitectura de microservicios contenerizados y mejorar la frecuencia de despliegue sin impactar la operación actual.
