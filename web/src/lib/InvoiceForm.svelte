<script lang="ts">
    import { browser } from "$app/environment";
    import { invoice } from "../store";

    const currencies = ["GBP (£)", "USD ($)", "EUR (€)"];
    const symbols = ["£", "$", "€"];
    const fromPlaceholder = `From Goldenhand Software
1 Electric Wharf, Generator Hall
Coventry, England, CV1 4JL
info@goldenhandosoftware.co.uk`;
    const toPlaceholder = `To Willy Wonka
Candy Factory, 1445 Norwood Ave
Itasca, IL 60143
willy@wonka.com
`;
    let filename: string = "";

    // total amount tracker
    $: if (browser) {
        $invoice.totalAmount = 0;
        for (const row of $invoice.rows) {
            if (row.rate == undefined || row.quantity == undefined) {
                row.amount = "";
            } else {
                row.rate = Math.floor(row.rate * 100) / 100;
                row.quantity = Math.floor(row.quantity * 100) / 100;
                row.amount = DecimalFixed(row.rate * row.quantity);
                $invoice.totalAmount += parseFloat(row.amount);
            }
        }
    }

    // currency tracker
    $: if (browser) {
        $invoice.symbol = CurrencySymbol($invoice.currency);
        for (const row of $invoice.rows) {
            row.symbol = CurrencySymbol($invoice.currency);
        }
        $invoice.rows = [...$invoice.rows];
    }

    function setFilename(event: Event) {
        const target = event.target as HTMLInputElement;
        if (target.files != undefined && target.files.length > 0) {
            filename = target.files[0].name;
        }
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
        $invoice.rows = [
            ...$invoice.rows,
            {
                description: "",
                symbol: $invoice.symbol,
                rate: undefined,
                quantity: undefined,
                amount: "",
            },
        ];
    }

    function clear() {
        $invoice.rows = [
            {
                description: "",
                symbol: symbols[0],
                rate: undefined,
                quantity: undefined,
                amount: "",
            },
        ];
        $invoice.currency = currencies[0];
        $invoice.symbol = symbols[0];
        $invoice.totalAmount = 0;
        $invoice.notes = "";
        $invoice.from = "";
        $invoice.to = "";
        $invoice.invoiceNumber = "";
        $invoice.issueDate = "";
        $invoice.dueDate = "";
    }

    async function submit(event: Event) {
        const target = event.target as HTMLFormElement;
        const formData = new FormData(target);

        // // deal with rows here
        formData.append("symbol", $invoice.symbol);
        formData.append("totalAmount", DecimalFixed($invoice.totalAmount));

        try {
            //const formData = new FormData(form);
            console.log("formData", formData);
            const response = await fetch("http://localhost:5555/generate/invoice", {
                method: "POST",
                body: formData,
            });

            console.log(response);
        } catch (error) {
            console.error(error);
        }
    }
</script>

<form
    class="container"
    on:submit|preventDefault={submit}
    method="POST"
    enctype="multipart/form-data">

    <div class="box box-good">
        <div class="columns">
            <div class="column">
                <p class="invoice-number-title"><b>Invoice Number</b></p>
                <input
                    class="input"
                    type="text"
                    name="invoiceNumber"
                    placeholder="Invoice 30"
                    bind:value={$invoice.invoiceNumber}
                    required
                />
            </div>
            <div class="column">
                <p class="logo-title"><b>Logo</b></p>
                <div class="file has-name is-fullwidth is-light">
                    <label class="file-label">
                        <input
                            class="file-input"
                            type="file"
                            name="logo"
                            
                            on:change={setFilename}
                        />
                        <span class="file-cta">
                            <span class="file-icon">
                                <i class="fa fa-upload" />
                            </span>
                            <span class="file-label"> Upload </span>
                        </span>
                        <span class="file-name"> {filename} </span>
                    </label>
                </div>
            </div>
            <div class="column">
                <p class="date-title"><b>Issue Date</b></p>
                <input
                    class="input"
                    type="date"
                    name="issueDate"
                    bind:value={$invoice.issueDate}
                    required
                />
            </div>
            <div class="column">
                <p class="date-title"><b>Due Date</b></p>
                <input
                    class="input"
                    type="date"
                    name="dueDate"
                    bind:value={$invoice.dueDate}
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
                        name="from"
                        bind:value={$invoice.from}
                    />
                </div>
                <div class="tile is-child">
                    <textarea
                        class="textarea"
                        placeholder={toPlaceholder}
                        name="to"
                        bind:value={$invoice.to}
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
                    {#each $invoice.rows as row}
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
                            Total: {$invoice.symbol}{DecimalFixed(
                                $invoice.totalAmount
                            )}
                        </p>
                    </li>
                </ul>
            </div>
        </div>

        <textarea
            class="textarea mt-5"
            placeholder="Optional Notes..."
            name="notes"
            rows="5"
            bind:value={$invoice.notes}
        />

        <div class="box-footer mt-5 has-text-right">
            <div class="select is-dark">
                <select bind:value={$invoice.currency}>
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
