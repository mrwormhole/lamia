export function DecimalFixed(n: number): string {
    return (Math.floor(n * 100) / 100).toFixed(2).toString();
}