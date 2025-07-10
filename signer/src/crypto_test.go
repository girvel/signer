package signer

import "testing"

const samplePrivateKey string = `-----BEGIN PRIVATE KEY-----
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
0WlfA4MPT+XwicTiFFoDCz2EgqHVx4M=
-----END PRIVATE KEY-----`

const sampleText string = `You will rejoice to hear that no disaster has accompanied the commencement of an enterprise which you have regarded with such evil forebodings. I arrived here yesterday, and my first task is to assure my dear sister of my welfare and increasing confidence in the success of my undertaking.
I am already far north of London, and as I walk in the streets of Petersburgh, I feel a cold northern breeze play upon my cheeks, which braces my nerves and fills me with delight. Do you understand this feeling? This breeze, which has travelled from the regions towards which I am advancing, gives me a foretaste of those icy climes. Inspirited by this wind of promise, my daydreams become more fervent and vivid. I try in vain to be persuaded that the pole is the seat of frost and desolation; it ever presents itself to my imagination as the region of beauty and delight. There, Margaret, the sun is for ever visible, its broad disk just skirting the horizon and diffusing a perpetual splendour. There—for with your leave, my sister, I will put some trust in preceding navigators—there snow and frost are banished; and, sailing over a calm sea, we may be wafted to a land surpassing in wonders and in beauty every region hitherto discovered on the habitable globe. Its productions and features may be without example, as the phenomena of the heavenly bodies undoubtedly are in those undiscovered solitudes. What may not be expected in a country of eternal light? I may there discover the wondrous power which attracts the needle and may regulate a thousand celestial observations that require only this voyage to render their seeming eccentricities consistent for ever. I shall satiate my ardent curiosity with the sight of a part of the world never before visited, and may tread a land never before imprinted by the foot of man. These are my enticements, and they are sufficient to conquer all fear of danger or death and to induce me to commence this laborious voyage with the joy a child feels when he embarks in a little boat, with his holiday mates, on an expedition of discovery up his native river. But supposing all these conjectures to be false, you cannot contest the inestimable benefit which I shall confer on all mankind, to the last generation, by discovering a passage near the pole to those countries, to reach which at present so many months are requisite; or by ascertaining the secret of the magnet, which, if at all possible, can only be effected by an undertaking such as mine.`

func TestSignThenVerify(t *testing.T) {
    cr, err := CreateCryptographerRSA([]byte(samplePrivateKey))

    if err != nil {
        t.Error(err.Error())
    }

    signature, err := cr.Sign(sampleText)

    if err != nil {
        t.Error(err.Error())
    }

    err = cr.Verify(sampleText, signature)

    if err != nil {
        t.Error(err.Error())
    }
}
