import React from "react"

type Highlight = {
    text: string
    mood: number
    category: string[]
    date: string
}

// ✅ Mood checker helper (optional improvement)
const moodChecker = (mood: number) => {
    switch (mood) {
        case 1:
            return "😞 Very Low"
        case 2:
            return "😐 Low"
        case 3:
            return "🙂 Neutral"
        case 4:
            return "😄 High"
        case 5:
            return "🔥 Very High"
        default:
            return "Unknown"
    }
}

// ✅ Card component using Highlight props
export const Card: React.FC<Highlight> = ({ text, mood, category, date }) => {
    return (
        <section className="p-4 rounded-lg border border-neutral-800 w-full">
            <h6 className="font-bold">{text}</h6>
            <p className="opacity-50 text-sm">{date} &#8226; {category.join(", ")}</p>
            <p>Mood: {moodChecker(mood)}</p>
        </section>
    )
}