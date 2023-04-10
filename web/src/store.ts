import { writable } from "svelte/store";
import { browser } from "$app/environment";

type InvoiceRow = {
    description: string;
    symbol: string;
    rate: number | undefined;
    quantity: number | undefined;
    amount: string;
}

type InvoiceData = {
    rows: Array<InvoiceRow>;
    currency: string;
    symbol: string;
    totalAmount: number;
    notes: string;
    from: string;
    to: string;
    invoiceNumber: string;
    issueDate: string;
    dueDate: string;
}

export const invoice = writable<InvoiceData>(fromLocalStorage());

invoice.subscribe(toLocalStorage);

// Get value from localStorage if in browser and the value is stored, otherwise fallback
function fromLocalStorage(key: string = "lamia-invoice-data"): InvoiceData {
    if (browser) {
        const stored = window.localStorage.getItem(key)

        if (stored != null) {
            return JSON.parse(stored);
        }
    }

    return {
        rows: [
            {
                description: "",
                symbol: "£",
                rate: undefined,
                quantity: undefined,
                amount: "",
            },
        ],
        currency: "GBP (£)",
        symbol: "£",
        totalAmount: 0,
        notes: "",
        from: "",
        to: "",
        invoiceNumber: "",
        issueDate: "",
        dueDate: "",
    }
}

// Set value to localStorage if in browser
function toLocalStorage(d: InvoiceData) {
    if (browser) {
        localStorage.setItem("lamia-invoice-data", JSON.stringify(d))
    }
}