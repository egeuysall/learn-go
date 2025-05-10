"use client"

import React, {useEffect, useState} from "react"
import {Card} from "../Card"

type Highlight = {
    text: string
    mood: number
    category: string[]
    date: string
    id: number
}

export default function Home() {
    const [highlights, setHighlights] = useState<Highlight[]>([])

    useEffect(() => {
        const fetchHighlights = async () => {
            const url = "http://134.122.15.244:8080/"

            try {
                const res = await fetch(url)
                if (!res.ok) throw new Error("Error fetching highlights.")

                const json = await res.json()
                setHighlights(json) // ✅ update state with fetched data
            } catch (err) {
                console.error("Error:", err)
            }
        }

        fetchHighlights()
    }, [])

    return (
      <main className="grid place-items-center h-screen w-screen">
        <section className="w-[90vw] grid gap-4">
          <h1 className="text-2xl md:text-3xl lg:text-4xl font-bold">Mood Tracker</h1>
          {highlights.length > 0 ? (
            <section className="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
              {highlights.map((highlight) => (
                <Card key={highlight.id} {...highlight} />
              ))}
            </section>
          ) : (
            <p className="text-neutral-500">No events could be found.</p>
          )}
        </section>
      </main>
    )
}