// English language
window.i18n = window.i18n || {};
window.i18n.en = {
    appTitle: 'Estimator',
    appTagline: 'Scientific* approach to task time estimation',
    appIrony: '* with a high degree of irony',
    initialEstimateTitle: 'Your initial time estimate',
    quantity: 'Quantity',
    timeUnit: 'Time unit',
    hours: 'Hours',
    days: 'Days',
    weeks: 'Weeks',
    months: 'Months',
    quarters: 'Quarters',
    pertParamsTitle: 'Additional parameters for PERT',
    optimisticEstimate: 'Optimistic estimate',
    mostLikelyEstimate: 'Most likely',
    pessimisticEstimate: 'Pessimistic estimate',
    selectMethodTitle: 'Select estimation method',
    piMethod: 'π Method (multiply by 3.14)',
    plus2Method: '"+2 units" Method',
    pertMethod: 'PERT Method',
    last10Method: '"Last 10%" Law',
    hofstadterMethod: 'Hofstadter\'s Law',
    combinedMethod: 'Combined (π + 2)',
    resultTitle: 'Estimation result',
    copyLink: 'Copy link',
    linkCopied: 'Link copied to clipboard',
    wisdomTitle: 'Project management wisdom',
    footer: 'Estimator © 2025 | Made with 💻 and 🤦‍♂️ | No scientific value',
    // Method descriptions
    methodDescriptions: {
        piMethodDesc: `
            <p>π Method: multiply your estimate by π (3.14159...).</p>
            <p class="mt-2">This method accounts for the natural tendency to underestimate task complexity and helps provide a more realistic forecast.</p>
        `,
        plus2MethodDesc: `
            <p>"+2 units" Method: add 2 units of measurement to your estimate.</p>
            <p class="mt-2">A simple and effective method that works well for iterative development and medium-sized tasks.</p>
        `,
        pertMethodDesc: function(o, m, p, result, sd) {
            return `
                <p>PERT Method: weighted average of three estimates: optimistic, most likely, and pessimistic.</p>
                <p class="mt-2">Formula: (O + 4M + P) / 6 = (${o} + 4×${m} + ${p}) / 6 = ${result}</p>
                <p class="mt-2">Standard deviation: ${sd}</p>
                <p class="mt-2">Probability ranges for completion:</p>
                <ul class="list-disc pl-5 mt-1">
                    <li>With 68% probability: ${result - sd} - ${result + sd}</li>
                    <li>With 95% probability: ${result - 2*sd} - ${result + 2*sd}</li>
                </ul>
            `;
        },
        last10MethodDesc: function(estimate, result) {
            return `
                <p>The "Last 10%" Law: The first 90% of the work takes 90% of the time, and the remaining 10% takes the other 90% of the time.</p>
                <p class="mt-2">This method adds another 90% to your initial estimate to account for unforeseen complexities in the final stage.</p>
                <p class="mt-2">Calculation: ${estimate} + (${estimate} × 0.9) = ${result}</p>
            `;
        },
        hofstadterMethodDesc: `
            <p>Hofstadter's Law: "It always takes longer than you expect, even when you take into account Hofstadter's Law."</p>
            <p class="mt-2">Doubles the estimate based on the principle "better safe than sorry".</p>
        `,
        combinedMethodDesc: function(estimate, piResult, result) {
            return `
                <p>Combined method: multiply by π, then add 2 units.</p>
                <p class="mt-2">For those who prefer to be maximally cautious. Use with care!</p>
                <p class="mt-2">Calculation: (${estimate} × π) + 2 = ${piResult} + 2 = ${result}</p>
            `;
        }
    },
    // Unit forms 
    unitForms: {
        hours: ['hour', 'hours', 'hours'],
        days: ['day', 'days', 'days'],
        weeks: ['week', 'weeks', 'weeks'],
        months: ['month', 'months', 'months'],
        quarters: ['quarter', 'quarters', 'quarters']
    },
    // Quotes
    quotes: [
        "The first 90% of the code accounts for the first 10% of the development time. The remaining 10% of the code accounts for the other 90% of the development time.",
        "Any task will take twice as long as you think it will take, even when you take this law into account.",
        "I don't know how to estimate, I'm just a programmer. — Every programmer during estimation",
        "If time estimation was an exact science, we would have finished this project already.",
        "The optimist believes we live in the best of all possible worlds. The pessimist fears this is true. The project manager knows it affects the timeline.",
        "Perfect code doesn't need estimation. It's already written.",
        "Hofstadter's Law: It always takes longer than you expect, even when you take into account Hofstadter's Law.",
        "Developer estimated the task would take 2 hours. Two days later: 'Almost done, just a bit more...'",
        "If a developer says 'it will take 5 more minutes', go get lunch.",
        "Q: How long does it take a programmer to finish a task? A: Just 5 more minutes, I'm almost done!"
    ]
};
