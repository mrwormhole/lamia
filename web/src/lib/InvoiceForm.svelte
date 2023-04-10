<script lang="ts">
    const currencies = ["GBP (£)", "USD ($)", "EUR (€)"];
    const symbols = ["£", "$", "€"];
    const fromPlaceholder = `From Goldenhand Software
1 Electric Wharf, Generator Hall
Coventry, England, CV1 4JL
info@goldenhandosoftware.co.uk`
    const toPlaceholder = `To Willy Wonka
Candy Factory, 1445 Norwood Ave
Itasca, IL 60143
willy@wonka.com
`

    let currency: string = currencies[0];
    let symbol: string = symbols[0];
    let totalAmount: number = 0;
    let notes: string = "";
    let from: string = "";
    let to: string = "";
    let invoiceNumber: string = "";
    let issueDate: string = "";
    let dueDate: string = "";

    type InvoiceRow = {
        description: string;
        symbol: string;
        rate: number | undefined;
        quantity: number | undefined;
        amount: string;
    };

    let rows: Array<InvoiceRow> = [
        {
            description: "",
            symbol: symbol,
            rate: undefined,
            quantity: undefined,
            amount: "",
        },
    ];

    // total amount calculator
    $: {
        totalAmount = 0;
        for (const row of rows) {
            if (row.rate == undefined || row.quantity == undefined) {
                row.amount = "";
            } else {
                row.rate = Math.floor(row.rate * 100) / 100;
                row.quantity = Math.floor(row.quantity * 100) / 100;
                row.amount = DecimalFixed(row.rate * row.quantity);
                totalAmount += parseFloat(row.amount);
            }
        }
    }

    // currency adjuster
    $: {
        symbol = CurrencySymbol(currency);
        for (const row of rows) {
            row.symbol = CurrencySymbol(currency);
        }
        rows = [...rows];
    }

    function CurrencySymbol(currency: string): string {
        for (const symbol of symbols) {
            if (currency.includes(symbol)) {
                return symbol;
            }
        }
        return "";
    }

    function setDecimal(event: Event) {
        const target = event.target as HTMLInputElement;
        target.value = DecimalFixed(parseFloat(target.value));
    }

    function DecimalFixed(n: number): string {
        return (Math.floor(n * 100) / 100).toFixed(2).toString();
    }

    function addRow() {
        rows = [
            ...rows,
            {
                description: "",
                symbol: symbol,
                rate: undefined,
                quantity: undefined,
                amount: "",
            },
        ];
    }

    function clear() {
        rows = [
            {
                description: "",
                symbol: symbol,
                rate: undefined,
                quantity: undefined,
                amount: "",
            },
        ];
        currency = currencies[0];
        symbol = symbols[0];
        totalAmount = 0;
    }

    function submit() {
        console.log("submitted: ", issueDate);
        console.log("submitted: ", dueDate);
    }
</script>

<form class="container" on:submit|preventDefault={submit}>
    <div class="box box-good">
        <div class="columns">
            <div class="column">
                <p class="invoice-number-title"><b>Invoice Number</b></p>
                <input
                    class="input"
                    type="text"
                    placeholder="Invoice 30"
                    bind:value={invoiceNumber}
                    required
                />
            </div>
            <div class="column">
                <p class="logo-title"><b>Logo</b></p>
                <div class="file has-name is-fullwidth is-light">
                    <label class="file-label">
                        <input class="file-input" type="file" name="logo" />
                        <span class="file-cta">
                            <span class="file-icon">
                                <i class="fa fa-upload" />
                            </span>
                            <span class="file-label"> Upload </span>
                        </span>
                        <span class="file-name"> logo.png </span>
                    </label>
                </div>
            </div>
            <div class="column">
                <p class="date-title"><b>Issue Date</b></p>
                <input
                    class="input"
                    type="date"
                    bind:value={issueDate}
                    required
                />
            </div>
            <div class="column">
                <p class="date-title"><b>Due Date</b></p>
                <input
                    class="input"
                    type="date"
                    bind:value={dueDate}
                    required
                />
            </div>
        </div>

        <div class="tile is-ancestor mt-5">
            <div class="tile is-parent is-12">
                <div class="tile is-child">
                    <textarea
                        class="textarea"
                        placeholder={fromPlaceholder}
                        bind:value={from}
                    />
                </div>
                <div class="tile is-child">
                    <textarea
                        class="textarea"
                        placeholder={toPlaceholder}
                        bind:value={to}
                    />
                </div>
            </div>
        </div>

        <div class="table-container mt-5">
            <table class="table is-fullwidth is-bordered">
                <thead>
                    <tr>
                        <th>Description</th>
                        <th>Rate</th>
                        <th>Quantity</th>
                        <th>Amount</th>
                    </tr>
                </thead>
                <tbody>
                    {#each rows as row}
                        <tr>
                            <td class="p-0">
                                <input
                                    id="description"
                                    class="input is-borderless"
                                    type="text"
                                    placeholder="Extreme Milky Chocolate High Quality 20% Cocoa"
                                    bind:value={row.description}
                                />
                            </td>
                            <td class="p-0">
                                <input
                                    id="rate"
                                    class="input is-borderless"
                                    type="number"
                                    on:change|preventDefault={setDecimal}
                                    min="0"
                                    step="1"
                                    placeholder="1.00"
                                    bind:value={row.rate}
                                />
                            </td>
                            <td class="p-0">
                                <input
                                    id="quantity"
                                    class="input is-borderless"
                                    type="number"
                                    on:change|preventDefault={setDecimal}
                                    min="0"
                                    step="1"
                                    placeholder="1.00"
                                    bind:value={row.quantity}
                                />
                            </td>
                            <td>
                                {row.symbol}{row.amount ? row.amount : "0.00"}
                            </td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>

        <div class="columns m-0">
            <div class="column is-offset-8 is-4 has-text-right">
                <ul class="list-group list-group-unbordered">
                    <li class="list-group-item">
                        <p class="subtitle has-text-weight-bold">
                            Total: {symbol}{DecimalFixed(totalAmount)}
                        </p>
                    </li>
                </ul>
            </div>
        </div>

        <textarea
            class="textarea mt-5"
            placeholder="Optional Notes..."
            rows="5"
            bind:value={notes}
        />

        <div class="box-footer mt-5 has-text-right">
            <div class="select is-dark">
                <select bind:value={currency}>
                    {#each currencies as currency}
                        <option>{currency}</option>
                    {/each}
                </select>
            </div>
            <button
                class="button is-light"
                type="button"
                on:click|preventDefault={addRow}
            >
                <span class="icon">
                    <i class="fa fa-plus-circle" aria-hidden="true" />
                </span>
                <span>Add a Row</span>
            </button>
            <button
                class="button is-light"
                type="button"
                on:click|preventDefault={clear}
            >
                <span class="icon">
                    <i class="fa fa-refresh" aria-hidden="true" />
                </span>
                <span>Clear</span>
            </button>
            <!--<button class="button is-light" on:click|preventDefault={clear}>
                <span class="icon">
                    <i class="fa fa-floppy-o" aria-hidden="true" />
                </span>
                <span>Save</span>
            </button>-->
            <button class="button is-light">
                <span class="icon">
                    <i class="fa fa-eye" aria-hidden="true" />
                </span>
                <span>Preview</span>
            </button>
            <button class="button is-dark" type="submit">
                <span class="icon">
                    <i class="fa fa-download" aria-hidden="true" />
                </span>
                <span>Download</span>
            </button>
        </div>
    </div>
</form>

<style>
    .invoice-number-title {
        font-size: 1rem;
        padding: 1rem 0rem;
    }

    .logo-title {
        font-size: 1rem;
        padding: 1rem 0rem;
    }

    .date-title {
        font-size: 1rem;
        padding: 1rem 0rem;
    }

    .is-borderless {
        border: 0;
        box-shadow: none;
    }

    .container {
        margin-top: 4rem;
        margin-bottom: 4rem;
    }
</style>
