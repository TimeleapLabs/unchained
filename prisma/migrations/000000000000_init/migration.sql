-- CreateTable
CREATE TABLE "AssetPrice" (
    "id" SERIAL NOT NULL,
    "dataSetId" INTEGER NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "block" INTEGER,
    "price" DECIMAL(65,30) NOT NULL,
    "signature" TEXT NOT NULL,

    CONSTRAINT "AssetPrice_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "DataSet" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "DataSet_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Signer" (
    "id" SERIAL NOT NULL,
    "key" TEXT NOT NULL,
    "name" TEXT,

    CONSTRAINT "Signer_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "SignersOnAssetPrice" (
    "signerId" INTEGER NOT NULL,
    "assetPriceId" INTEGER NOT NULL,

    CONSTRAINT "SignersOnAssetPrice_pkey" PRIMARY KEY ("signerId","assetPriceId")
);

-- CreateIndex
CREATE UNIQUE INDEX "AssetPrice_dataSetId_block_key" ON "AssetPrice"("dataSetId", "block");

-- CreateIndex
CREATE UNIQUE INDEX "DataSet_name_key" ON "DataSet"("name");

-- CreateIndex
CREATE UNIQUE INDEX "Signer_key_key" ON "Signer"("key");

-- AddForeignKey
ALTER TABLE "AssetPrice" ADD CONSTRAINT "AssetPrice_dataSetId_fkey" FOREIGN KEY ("dataSetId") REFERENCES "DataSet"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "SignersOnAssetPrice" ADD CONSTRAINT "SignersOnAssetPrice_assetPriceId_fkey" FOREIGN KEY ("assetPriceId") REFERENCES "AssetPrice"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "SignersOnAssetPrice" ADD CONSTRAINT "SignersOnAssetPrice_signerId_fkey" FOREIGN KEY ("signerId") REFERENCES "Signer"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

