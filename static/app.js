// app.js

// Helper to call the integral endpoint
async function computeIntegral(expression, lower, upper, steps) {
    const url = `/integral?expression=${encodeURIComponent(expression)}&lower_bound=${lower}&upper_bound=${upper}&steps=${steps}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to compute integral");
    }
    return await res.json();
}

// Helper to call the limit endpoint
async function computeLimit(expression, point) {
    const url = `/limit?expression=${encodeURIComponent(expression)}&point=${point}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to compute limit");
    }
    return await res.json();
}

async function computeDerivative(expression, point) {
    const url = `/derivative?expression=${encodeURIComponent(expression)}&point=${point}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to compute derivative");
    }
    return await res.json();
}

async function computeCartesianToPolar(x, y, z) {
    const url = `/cartesian-to-polar?x=${x}&y=${y}&z=${z}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to convert to polar coordinates");
    }
    return await res.json();
}

async function computeCartesianToSpherical(x, y, z) {
    const url = `/cartesian-to-spherical?x=${x}&y=${y}&z=${z}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to convert to spherical coordinates");
    }
    return await res.json();
}

async function computePolarToCartesian(r, theta, z) {
    const url = `/polar-to-cartesian?r=${r}&theta=${theta}&z=${z}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to convert to cartesian coordinates");
    }
    return await res.json();
}

async function computePolarToSpherical(r, theta, z) {
    const url = `/polar-to-spherical?r=${r}&theta=${theta}&z=${z}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to convert to spherical coordinates");
    }
    return await res.json();
}

async function computeSphericalToCartesian(r, theta, phi) {
    const url = `/spherical-to-cartesian?r=${r}&theta=${theta}&phi=${phi}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to convert to cartesian coordinates");
    }
    return await res.json();
}

async function computeSphericalToPolar(r, theta, phi) {
    const url = `/spherical-to-polar?r=${r}&theta=${theta}&phi=${phi}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to convert to polar coordinates");
    }
    return await res.json();
}

async function computeCircleArea(radius) {
    const url = `/area-circle?radius=${radius}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to compute circle area");
    }
    return await res.json();
}

async function computeCircleCircumference(radius) {
    const url = `/circumference-circle?radius=${radius}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to compute circle circumference");
    }
    return await res.json();
}

async function computeRectangleArea(length, width) {
    const url = `/area-rectangle?length=${length}&width=${width}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to compute rectangle area");
    }
    return await res.json();
}

async function computeRectanglePerimeter(length, width) {
    const url = `/perimeter-rectangle?length=${length}&width=${width}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to compute rectangle perimeter");
    }
    return await res.json();
}

async function grabPlotData(expression, lower, upper) {
    const url = `/plotdata?expression=${encodeURIComponent(expression)}&lower_bound=${lower}&upper_bound=${upper}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to fetch plot data");
    }
    return await res.json();

}

async function grabTangentLineData(expression, point, lower, upper) {
    const url = `/tangentline?expression=${encodeURIComponent(expression)}&point=${point}&lower_bound=${lower}&upper_bound=${upper}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to fetch plot data");
    }
    return await res.json();

}

async function grabPointData(expression, point, lower, upper) {
    const url = `/evaluate2Dfunctionpoint?expression=${encodeURIComponent(expression)}&point=${point}`;
    const res = await fetch(url);
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || "Failed to fetch plot data");
    }
    return await res.json();

}

async function createIntegralPlotParts(expression, lower, upper) {
    const plot_data = await grabPlotData(expression, lower, upper);
    const [x, y] = plot_data["plot_data"]

    const data = {
        x: x,
        y: y,
        type: 'scatter',
        fill: 'tozeroy',
    }
    const layout = {
        title: 'Integral Plot',
        xaxis: {
            title: 'X-axis'
        },
        yaxis: {
            title: 'Y-axis'
        }
    };
    return { data, layout };
}

async function createLimitPlotParts(expression, point, limit) {
    const lower = point - 1;
    const upper = point + 1;
    const plot_data = await grabPlotData(expression, lower, upper);
    const [x, y] = plot_data["plot_data"]

    const data = [{
        x: x,
        y: y,
        type: 'scatter',
        name: 'Function Plot'
    },
    {
        x: [point],
        y: [limit],
        type: 'scatter',
        mode: 'markers',
        line: { color: 'red' },
        name: 'Limit Line'
    }
    ]
    const layout = {
        title: 'Limit Plot',
        xaxis: {
            title: 'X-axis'
        },
        yaxis: {
            title: 'Y-axis'
        }
    };
    return { data, layout };
}

async function createDerivativePlotParts(expression, point) {
    const lower = point - 10;
    const upper = point + 10;
    const plot_data = await grabPlotData(expression, lower, upper);
    const [x, y] = plot_data["plot_data"]
    const tangent_data = await grabTangentLineData(expression, point, lower, upper);
    const [tx, ty] = tangent_data["tangent_line"];
    const point_data = await grabPointData(expression, point);
    const [tpx, tpy] = point_data["point_data"];
    console.log(tpx, tpy);

    const data = [{
        x: x,
        y: y,
        type: 'scatter',
        name: 'Function Plot'
    },
    {
        x: tx,
        y: ty,
        type: 'scatter',
        mode: 'lines',
        line: { color: 'red' },
        name: 'Tangent Line'
    },
    {
        x: [tpx],
        y: [tpy],
        type: 'scatter',
        mode: 'markers',
        line: { color: 'green' },
        name: 'Tangent Point'
    }
    ]
    const layout = {
        title: 'Derivative Plot',
        xaxis: {
            title: 'X-axis'
        },
        yaxis: {
            title: 'Y-axis'
        }
    };
    return { data, layout };
}

// Expose functions to global scope (optional for simplicity)
window.MathAPI = {
    computeIntegral,
    computeLimit,
    computeDerivative,
    computeCartesianToPolar,
    computeCartesianToSpherical,
    computePolarToCartesian,
    computePolarToSpherical,
    computeSphericalToCartesian,
    computeSphericalToPolar,
    computeCircleArea,
    computeCircleCircumference,
    computeRectangleArea,
    computeRectanglePerimeter,
    createIntegralPlotParts,
    createLimitPlotParts,
    createDerivativePlotParts,
};
