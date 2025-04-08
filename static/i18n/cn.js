// Chinese language
window.i18n = window.i18n || {};
window.i18n.cn = {
    appTitle: '时间估算器',
    appTagline: '任务时间估算的科学*方法',
    appIrony: '* 带有高度讽刺意味',
    initialEstimateTitle: '您的初始时间估计',
    quantity: '数量',
    timeUnit: '时间单位',
    hours: '小时',
    days: '天',
    weeks: '周',
    months: '月',
    quarters: '季度',
    pertParamsTitle: 'PERT方法的附加参数',
    optimisticEstimate: '乐观估计',
    mostLikelyEstimate: '最可能',
    pessimisticEstimate: '悲观估计',
    selectMethodTitle: '选择估算方法',
    piMethod: 'π方法 (乘以3.14)',
    plus2Method: '"+2单位"方法',
    pertMethod: 'PERT方法',
    last10Method: '"最后10%"法则',
    hofstadterMethod: '霍夫施塔特定律',
    combinedMethod: '组合方法 (π + 2)',
    resultTitle: '估算结果',
    copyLink: '复制链接',
    linkCopied: '链接已复制到剪贴板',
    wisdomTitle: '项目管理的智慧',
    footer: '时间估算器 © 2025 | 用💻和🤦‍♂️制作 | 没有科学价值',
    // Method descriptions
    methodDescriptions: {
        piMethodDesc: `
            <p>π方法：将您的估计乘以π (3.14159...)。</p>
            <p class="mt-2">此方法考虑了低估任务复杂性的自然倾向，有助于提供更现实的预测。</p>
        `,
        plus2MethodDesc: `
            <p>"+2单位"方法：在您的估计上增加2个测量单位。</p>
            <p class="mt-2">一种简单有效的方法，适用于迭代开发和中等规模任务。</p>
        `,
        pertMethodDesc: function(o, m, p, result, sd) {
            return `
                <p>PERT方法：三种估计的加权平均：乐观的、最可能的和悲观的。</p>
                <p class="mt-2">公式：(O + 4M + P) / 6 = (${o} + 4×${m} + ${p}) / 6 = ${result}</p>
                <p class="mt-2">标准差：${sd}</p>
                <p class="mt-2">完成的概率范围：</p>
                <ul class="list-disc pl-5 mt-1">
                    <li>68%的概率：${result - sd} - ${result + sd}</li>
                    <li>95%的概率：${result - 2*sd} - ${result + 2*sd}</li>
                </ul>
            `;
        },
        last10MethodDesc: function(estimate, result) {
            return `
                <p>"最后10%"法则：工作的前90%占用90%的时间，剩下的10%又占用另外90%的时间。</p>
                <p class="mt-2">此方法在初始估计的基础上再增加90%，以应对最终阶段可能出现的不可预见的复杂性。</p>
                <p class="mt-2">计算：${estimate} + (${estimate} × 0.9) = ${result}</p>
            `;
        },
        hofstadterMethodDesc: `
            <p>霍夫施塔特定律："做事总是比你预期的要花更长的时间，即使你考虑到了霍夫施塔特定律。"</p>
            <p class="mt-2">基于"宁可过度准备也不要错过期限"的原则，将估计值翻倍。</p>
        `,
        combinedMethodDesc: function(estimate, piResult, result) {
            return `
                <p>组合方法：先乘以π，然后加2个单位。</p>
                <p class="mt-2">适合那些喜欢最大程度谨慎的人。谨慎使用！</p>
                <p class="mt-2">计算：(${estimate} × π) + 2 = ${piResult} + 2 = ${result}</p>
            `;
        }
    },
    // Unit forms
    unitForms: {
        hours: ['小时', '小时', '小时'],
        days: ['天', '天', '天'],
        weeks: ['周', '周', '周'],
        months: ['月', '月', '月'],
        quarters: ['季度', '季度', '季度']
    },
    // Quotes
    quotes: [
        "代码的前90%占用开发时间的10%。剩下的10%代码占用了剩余90%的开发时间。",
        "任何任务都会比你想象的花费两倍的时间，即使你已经考虑到了这个规律。",
        "我不知道如何估算，我只是一个程序员。 — 每个程序员在估算时",
        "如果时间估算是一门精确的科学，我们这个项目早就完成了。",
        "乐观主义者相信我们生活在所有可能世界中最好的一个。悲观主义者担心这是真的。项目经理知道这会影响时间表。",
        "完美的代码不需要估算。它已经被写好了。",
        "霍夫施塔特定律：做事总是比你预期的要花更长的时间，即使你考虑到了霍夫施塔特定律。",
        "开发者估计任务需要2小时。两天后：'快完成了，再差一点点...'",
        "如果开发者说'还需要5分钟'，那就去吃午饭吧。",
        "问：程序员需要多长时间才能完成一项任务？答：再5分钟，我快完成了！"
    ]
};
