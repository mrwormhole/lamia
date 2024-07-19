import { browser } from "$app/environment";
import { writable } from "svelte/store";

type InvoiceRow = {
    description: string;
    rate: number | undefined;
    quantity: number | undefined;
    amount: string;
};

type InvoiceData = {
    rows: Array<InvoiceRow>;
    currencySymbol: string;
    totalAmount: number;

    notes: string;
    from: string;
    to: string;
    invoiceNumber: string;
    issueDate: string;
    dueDate: string;
    logoFilename: string;
    logoBase64Img: string;
};

export const invoice = writable<InvoiceData>(fromLocalStorage());
invoice.subscribe(toLocalStorage);

// Get the values from localStorage if in browser and the value is stored
function fromLocalStorage(key = "lamia-invoice-data"): InvoiceData {
    const empty: InvoiceData = {
        rows: [
            {
                description: "",
                rate: undefined,
                quantity: undefined,
                amount: "",
            },
        ],
        currencySymbol: "£",
        totalAmount: 0,
        notes: "",
        from: "",
        to: "",
        invoiceNumber: "",
        issueDate: "",
        dueDate: "",
        logoFilename: "",
        logoBase64Img: "",
    };
    if (!browser) {
        return empty;
    }

    const stored = window.localStorage.getItem(key);
    if (stored != null) {
        return JSON.parse(stored);
    }

    return empty;
}

// Set the values to localStorage if in browser
function toLocalStorage(d: InvoiceData) {
    if (!browser) {
        return;
    }
    localStorage.setItem("lamia-invoice-data", JSON.stringify(d));
}
