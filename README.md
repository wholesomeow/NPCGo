# NPC GO

NPCGo is an improved NPC Generator to be used for generating more detailed NPCs for your games. Either adhoc when your party decides they want to talk to some random person in a crowd that you didn't plan for or when you need help creating characters that feel as real as the players themselves.

The tool is written in golang and uses grammar-based rule sets to turn a collection of words into a description worthy of being read at your table.

### Basic features of the NPC:

- Physical Description
- Sex, Gender, and Sexual Orientation (for the Bards out there)
- Job, Skills, and Stat Block
- Race, Name, and Pronouns

### Personality features of the NPC:

- Cognitive Science domains
- Enneagram description
- MICE description
- Various personality frameworks

The personality features all work together to create a more wholistic representation of a person, by giving them fears, hopes, desires, attachement styles, mental health descriptions and more.

As the project is developed more, the amount of information will be expanded on as well as become both more targeted and robust. Hopefully GMs will be able to use this tool to develop very robust NPCs on the fly that can provide a solid foundation for them to focus on the storytelling experience for their parties.

## How It Works

When a user presses the "Generate" button, the basic information is first collected from several datasets to generate a base for the NPC. These are things like Race, Subrace, Gender, Sex, Sexual Orientation, and their Job.

Along with their Job are things like Job Status and Societal Status to help give the GM and indication of the NPCs status and reputation in the world. **This is not yet fully implemented**

After that information is compiled, the personality data is compiled. This tool uses a few personality test frameworks to provide a sort of baseline that is grounded in reality. The primary framework being the [Enneagram Personality](https://www.enneagraminstitute.com/how-the-enneagram-system-works) framwork. It randomly selects one of the 9 personality architypes and then selects that types "Wing", along with a Blend percentage to create some nuance in the personality type created. For more information on why that framework was chosen, read the Medium article here. **Article coming soon**

Once the NPC dataset is compiled, it is referenced against local databases and WordNets to greatly expand the descriptors available in the NPC dataset. After that referencing is completed, it is then pushed through the grammar-based rule sets to turn the collection of words into sentences, repeating the process growing more complex each time. After a set amount of iterations, the generation is QA checked internally and then displayed on screen. No need for AI.

Nice and easy.

# Install

TBD

# Usage

TBD

## Acknowledgments

- https://github.com/miethe for their [DnD Character Generator](https://github.com/miethe/DnD-Character-Generator) that inspired some of the early design I had, provided me a starting point for the data I generated, and giving me their dice rolling script so I didn't have to labour over doing the math myself.
- https://github.com/bvezilic for their [DnD Name Generator](https://github.com/bvezilic/DnD-name-generator) that also served as a great point of inspiration and guidance for some of the early design.
- The Enneagram Institue for providing a great amount of data to start with.
- John Trubys _The Anatomy of Story_ for giving me a great framework to creating interesting characters for interesting storytelling. Please go buy the book.
- My buddy Phil for helping me figure out how best to make this project happen and avoid all the rabbit holes and pitfalls I really wanted to throw myself down.
