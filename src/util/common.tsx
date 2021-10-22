export function Epoch2Duration(epoch: number): string {
  return (
    Math.floor(epoch / 86400) +
    "d " +
    new Date((epoch % 86400) * 1000)
      .toUTCString()
      .replace(/.*(\d{2}):(\d{2}):(\d{2}).*/, "$1h $2m $3s")
  );
}
