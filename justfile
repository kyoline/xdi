default:
   @just --list --list-heading $'available recipes:\n' --list-prefix $'   -> '

run:
   deno run --allow-write --allow-read main.ts