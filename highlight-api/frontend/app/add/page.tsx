"use client"

import React, {useState} from "react"

export default function NewHighlightPage() {
    const [form, setForm] = useState({
        text: "",
        mood: 3,
        category: "",
        date: ""
    })

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target
        setForm((prev) => ({
            ...prev,
            [name]: name === "mood" ? Number(value) : value
        }))
    }

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()

        // Fix: sending a single object instead of an array
        const payload = {
            ...form,
            category: form.category.split(",").map((c) => c.trim())
        }

        try {
            const res = await fetch("http://localhost:8080/highlight", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(payload)
            })

            if (!res.ok) {
                throw new Error("Failed to submit highlight")
            }
            alert("Highlight submitted!")
            setForm({ text: "", mood: 3, category: "", date: "" })
        } catch (err) {
            console.error("Error submitting: ", err) // Improved logging
            alert("Error submitting: " + err)
        }
    }

    return (
        <main className="grid place-items-center h-screen w-screen">
            <section className="grid gap-4 w-[90vw] max-w-xl">
                <form onSubmit={handleSubmit} className="max-w-xl mx-auto p-6 space-y-4">
                    <input
                        name="text"
                        placeholder="Highlight text"
                        value={form.text}
                        onChange={handleChange}
                        className="w-full p-2.5 border border-neutral-800 rounded-lg"
                        required
                    />
                    <input
                        name="date"
                        type="date"
                        value={form.date}
                        onChange={handleChange}
                        className="w-full p-2.5 border border-neutral-800 rounded-lg"
                        required
                    />
                    <input
                        name="category"
                        placeholder="Comma-separated categories"
                        value={form.category}
                        onChange={handleChange}
                        className="w-full p-2.5 border border-neutral-800 rounded-lg"
                        required
                    />
                    <select
                        name="mood"
                        value={form.mood}
                        onChange={handleChange}
                        className="w-full p-2.5 border border-neutral-800 rounded-lg appearance-none"
                    >
                        <option value={1}>😞 Very Low</option>
                        <option value={2}>😐 Low</option>
                        <option value={3}>🙂 Neutral</option>
                        <option value={4}>😄 High</option>
                        <option value={5}>🔥 Very High</option>
                    </select>
                    <button type="submit" className="w-full p-2.5 border border-neutral-800 rounded-lg">
                        Submit Highlight
                    </button>
                </form>
            </section>
        </main>
    )
}