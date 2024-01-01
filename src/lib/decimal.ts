export function DecimalFixed(n: number): string {
    return DecimalFixedNum(n).toFixed(2);
}

export function DecimalFixedNum(n: number) {
    let res: number = parseFloat((n * 100).toFixed(2));
    res = Math.floor(res) / 100;
    return res;
}
