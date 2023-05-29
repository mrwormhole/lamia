<script lang="ts">
    import { browser } from "$app/environment";
    import { goto } from "$app/navigation";
    import { invoice } from "../store";
    import { DecimalFixed } from "./decimal";

    const currencies = ["£", "$", "€"];
    const fromPlaceholder = `From Goldenhand Software
1 Electric Wharf, Generator Hall
Coventry, England, CV1 4JL
info@goldenhandosoftware.co.uk`;
    const toPlaceholder = `To Willy Wonka
Candy Factory, 1445 Norwood Ave
Itasca, IL 60143
willy@wonka.com
`;
    const validImageTypes = ["image/jpg", "image/jpeg", "image/png"];
    const mb = 1 << (10 * 2);

    let notifications: Array<string> = [];

    $: if (browser) {
        // total amount tracker
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

    function setLogo(event: Event) {
        const target = event.target as HTMLInputElement;
        if (target.files != undefined && target.files.length > 0) {
            const file = target.files[0];

            if (!validImageTypes.includes(file.type)) {
                notifications = [
                    ...notifications,
                    `Attached logo(${file.type}) must be an image file.`,
                ];
                return;
            }

            if (file.size > 4 * mb) {
                notifications = [
                    ...notifications,
                    `Attached logo size(${file.size} bytes) must be less than 4MB.`,
                ];
                return;
            }

            const reader: FileReader = new FileReader();
            reader.onload = function () {
                if (reader.result != undefined) {
                    $invoice.logoBase64Img = reader.result.toString();
                    $invoice.logoFilename = file.name;

                    console.log("SIZE", file.size);
                }
            };
            reader.onerror = function (error) {
                notifications = [...notifications, `Error occurred, ${error}.`];
            };

            reader.readAsDataURL(file);
        }
    }

    function setDecimal(event: Event) {
        const target = event.target as HTMLInputElement;
        target.value = DecimalFixed(parseFloat(target.value));
    }

    function addRow() {
        $invoice.rows = [
            ...$invoice.rows,
            {
                description: "",
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
                rate: undefined,
                quantity: undefined,
                amount: "",
            },
        ];
        $invoice.currencySymbol = currencies[0];
        $invoice.totalAmount = 0;
        $invoice.notes = "";
        $invoice.from = "";
        $invoice.to = "";
        $invoice.invoiceNumber = "";
        $invoice.issueDate = "";
        $invoice.dueDate = "";
        $invoice.logoBase64Img = "";
        $invoice.logoFilename = "";
    }
</script>

<form
    class="container"
    on:submit|preventDefault={() => {
        notifications = [];
        goto("/print");
    }}
>
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
                            accept={validImageTypes.join(",")}
                            on:change={setLogo}
                        />
                        <span class="file-cta">
                            <span class="file-icon">
                                <i class="fa fa-upload" />
                            </span>
                            <span class="file-label"> Upload </span>
                        </span>
                        <span class="file-name">
                            {$invoice.logoFilename ? $invoice.logoFilename : ""}
                        </span>
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
                        <th>Price</th>
                        <th>Qty</th>
                        <th>Total</th>
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
                                    required
                                />
                            </td>
                            <td class="p-0">
                                <input
                                    id="rate"
                                    class="input is-borderless"
                                    type="number"
                                    on:change|preventDefault={setDecimal}
                                    min="0"
                                    step="0.01"
                                    placeholder="1.00"
                                    bind:value={row.rate}
                                    required
                                />
                            </td>
                            <td class="p-0">
                                <input
                                    id="quantity"
                                    class="input is-borderless"
                                    type="number"
                                    on:change|preventDefault={setDecimal}
                                    min="0"
                                    step="0.01"
                                    placeholder="1.00"
                                    bind:value={row.quantity}
                                    required
                                />
                            </td>
                            <td>
                                {$invoice.currencySymbol}{row.amount
                                    ? row.amount
                                    : "0.00"}
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
                            Total: {$invoice.currencySymbol}{DecimalFixed(
                                $invoice.totalAmount
                            )}
                        </p>
                    </li>
                </ul>
            </div>
        </div>

        <textarea
            class="textarea mt-5 mb-5"
            placeholder="Optional Notes..."
            name="notes"
            rows="12"
            bind:value={$invoice.notes}
        />

        {#if notifications.length != 0}
            <div class="notification is-danger">
                <button
                    class="delete"
                    on:click|preventDefault={() => (notifications = [])}
                />
                <ul class="">
                    {#each notifications as notification}
                        <li class="">
                            - {notification}
                        </li>
                    {/each}
                </ul>
            </div>
        {/if}

        <div class="box-footer mt-5 has-text-right">
            <div class="select is-dark">
                <select bind:value={$invoice.currencySymbol}>
                    {#each ["GBP (£)", "USD ($)", "EUR (€)"] as currencyText, i}
                        <option value={currencies[i]}>{currencyText}</option>
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
                    <i class="fa fa-print" aria-hidden="true" />
                </span>
                <span>Print</span>
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
