package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
другом   Кристофером   Робином,   головой   вниз,  пересчитывая
ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
	кажется, что можно бы найти какой-то другой способ, если бы  он
только   мог   на  минутку  перестать  бумкать  и  как  следует
сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
	Как бы то ни было, вот он уже спустился  и  готов  с  вами
познакомиться.
- Винни-Пух. Очень приятно!
	Вас,  вероятно,  удивляет, почему его так странно зовут, а
если вы знаете английский, то вы удивитесь еще больше.
	Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
можешь  сделать вид, что ты просто понарошку стрелял; а если ты
звал его тихо, то все подумают, что ты  просто  подул  себе  на
нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
Робин решил отдать его своему медвежонку, чтобы оно не  пропало
зря.
	А  Винни - так звали самую лучшую, самую добрую медведицу
в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
забыл.
	Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
	Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
иногда,  особенно  когда  папа  дома,  он больше любит тихонько
посидеть у огня и послушать какую-нибудь интересную сказку.
	В этот вечер...
	As you can see, he goes down the stairs after his
friend Christopher Robin, head down, counting
the steps with the back of his head: boom-boom-boom.  
He doesn't know any other way to get off the stairs yet.  Sometimes to him, really,
	it seems that some other way could be found if he
could just stop booming for a minute and
concentrate properly. But alas - he has no time to concentrate.
	Anyway, here he is already down and ready to
meet you.
- Winnie the Pooh. Very nice!
	You are probably wondering why his name is so strange, and
if you know English, then you will be even more surprised.
	This extraordinary name was given to him by Christopher Robin.  Need
I must tell you that Christopher Robin once knew a
swan on the pond, whom he called Pooh. It was
a very appropriate name for a swan, because if you call a swan
loudly: "Pooh! Pooh!"- and he does not respond, then you always
you can pretend that you were just pretending to shoot; and if you
called him softly, then everyone will think that you just blew on yourself
nose.  The swan then disappeared somewhere, but the name remained, and Christopher
Robin decided to give it to his little bear so that it would not
be wasted.
	And Winnie was the name of the best, kindest bear
in the zoological garden, which Christopher really, really loved
Robin.  And she really, really loved him. Whether she was named Winnie
after Pooh, or Pooh was named after her - no
one knows now, not even Christopher Robin's dad. Once he knew, but now
he has forgotten.
	In short, now the bear's name is Winnie the Pooh, and you know why.
	Sometimes Winnie-the-Pooh likes to play something in the evening, and
sometimes, especially when Dad is at home, he likes to be quiet more
sit by the fire and listen to some interesting fairy tale.
	This evening...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"the",         // 14
				"to",          // 11
				"you",         // 11
				"and",         // 9
				"he",          // 9
				"он",          // 8
				"-",           // 7
				"Christopher", // 6
				"name",        // 6
				"а",           // 6
				"и",           // 6
				"a",           // 5
				"that",        // 5
				"was",         // 5
				"ты",          // 5
				"что",         // 5
				"Winnie",      // 4
				"be",          // 4
				"his",         // 4
				"if",          // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}
